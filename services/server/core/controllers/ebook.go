package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	oss "github.com/ilovelili/aliyun-client/oss"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

const (
	// chromeDevTool chrome headless devtool endpoint
	chromeDevTool = "http://127.0.0.1:9222"

	// pdf width in inches
	pdfWidth float64 = 8.27

	// pdf height in inches
	pdfHeight float64 = 11.64
)

// EbookController ebook controller
type EbookController struct {
	repository *repositories.EbookRepository
	svc        *oss.Service
}

// NewEbookController new controller
func NewEbookController() *EbookController {
	config := utils.GetConfig()
	return &EbookController{
		repository: repositories.NewEbookRepository(),
		svc: func() *oss.Service {
			_svc := oss.NewService(config.OSS.APIKey, config.OSS.APISecret)
			_svc.SetEndPoint(config.OSS.Endpoint)
			_svc.SetBucket(config.OSS.BucketName)
			return _svc
		}(),
	}
}

// SaveEbook save ebook
func (c *EbookController) SaveEbook(ebook *models.Ebook) error {
	ebook.ResolveHash()
	dirty, err := c.repository.Upsert(ebook)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveEbook)
	}

	// if dirty, and then
	if dirty {
		// 1. upload to storage
		if err = c.uploadToStorage(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToUploadEbookToCloud)
		}

		// 2. generate pdf / img file thru chrome headless
		if err = c.convertHTML(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToConvertEbookHTML)
		}
	}

	return nil
}

// uploadToCloudStorage upload css folder and index.html to aliyun
// TODO: clear local file storage when domain gets ready and can be hosted by aliyun oss
func (c *EbookController) uploadToStorage(ebook *models.Ebook) error {
	// step 1. create corresponding directory (css / html)
	cssdiropts := &oss.UploadOptions{
		Public:       true,
		ObjectName:   ebook.Date,
		ParentFolder: fmt.Sprintf("ebook/css/%s/%s/%s/", ebook.Year, ebook.Class, ebook.Name),
		IsFolder:     true,
	}
	cssdirrespchan := c.svc.AsyncUpload(cssdiropts)

	htmldiropts := &oss.UploadOptions{
		Public:       true,
		ObjectName:   ebook.Date,
		ParentFolder: fmt.Sprintf("ebook/html/%s/%s/%s/", ebook.Year, ebook.Class, ebook.Name),
		IsFolder:     true,
	}
	htmldirrespchan := c.svc.AsyncUpload(htmldiropts)

	pwd, _ := os.Getwd()
	htmllocaldir := path.Join(pwd, "ebook", ebook.Year, ebook.Class, ebook.Name, ebook.Date)
	csslocaldir := path.Join(htmllocaldir, "css")

	_, err := os.Stat(csslocaldir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(csslocaldir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var (
		cssfilerespchan  chan *oss.UploadResponse
		htmlfilerespchan chan *oss.UploadResponse
	)

	// step 2. upload css file
	if cssdirresp := <-cssdirrespchan; cssdirresp.Error != nil {
		return cssdirresp.Error
	}

	csslocalfile := path.Join(csslocaldir, "style.css")
	err = ioutil.WriteFile(csslocalfile, []byte(ebook.ResolveCloudCSS()), os.ModePerm)
	// defer os.Remove(csspath)
	if err != nil {
		return err
	}

	cssfileopts := &oss.UploadOptions{
		Public:       true,
		ObjectName:   csslocalfile,
		ParentFolder: fmt.Sprintf("ebook/css/%s/%s/%s/%s", ebook.Year, ebook.Class, ebook.Name, ebook.Date),
	}
	cssfilerespchan = c.svc.AsyncUpload(cssfileopts)

	// step 3. upload html file
	if htmldirresp := <-htmldirrespchan; htmldirresp.Error != nil {
		return htmldirresp.Error
	}

	htmllocalfile := path.Join(htmllocaldir, "index.html")
	err = ioutil.WriteFile(htmllocalfile, []byte(ebook.ResolveCloudHTML()), os.ModePerm)
	// defer os.Remove(htmllocalfile)
	if err != nil {
		return err
	}

	htmlfileopts := &oss.UploadOptions{
		Public:       true,
		ObjectName:   htmllocalfile,
		ParentFolder: fmt.Sprintf("ebook/html/%s/%s/%s/%s", ebook.Year, ebook.Class, ebook.Name, ebook.Date),
	}
	htmlfilerespchan = c.svc.AsyncUpload(htmlfileopts)

	// wait for upload
	if cssfileresp := <-cssfilerespchan; cssfileresp.Error != nil {
		return cssfileresp.Error
	}

	if htmlfileresp := <-htmlfilerespchan; htmlfileresp.Error != nil {
		return htmlfileresp.Error
	}

	return nil
}

// convertHTML convert ebook html to pdf and jpg
func (c *EbookController) convertHTML(ebook *models.Ebook) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New(chromeDevTool)
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return
		}
	}
	defer devt.Close(ctx, pt)

	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return
	}
	defer conn.Close() // Leaving connections open will leak memory.

	cli := cdp.NewClient(conn)
	// Open a DOMContentEventFired client to buffer this event.
	domContent, err := cli.Page.DOMContentEventFired(ctx)
	if err != nil {
		return
	}
	defer domContent.Close()

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = cli.Page.Enable(ctx); err != nil {
		return
	}

	pwd, _ := os.Getwd()
	htmllocaldir := path.Join(pwd, "ebook", ebook.Year, ebook.Class, ebook.Name, ebook.Date)
	// Create the Navigate arguments
	navArgs := page.NewNavigateArgs(fmt.Sprintf("file://%s", path.Join(htmllocaldir, "index.html")))
	nav, err := cli.Page.Navigate(ctx, navArgs)
	if err != nil {
		return
	}

	// Wait until we have a DOMContentEventFired event.
	if _, err = domContent.Recv(); err != nil {
		return
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	imgOutput := path.Join(htmllocaldir, "output.jpg")
	// Capture a screenshot of the current page.
	screenshotArgs := page.NewCaptureScreenshotArgs().
		SetFormat("jpeg").
		SetQuality(100)

	screenshot, err := cli.Page.CaptureScreenshot(ctx, screenshotArgs)
	if err != nil {
		return
	}
	if err = ioutil.WriteFile(imgOutput, screenshot.Data, 0644); err != nil {
		return
	}

	fmt.Printf("Saved screenshot: %s\n", imgOutput)

	// Print to PDF
	printToPDFArgs := page.NewPrintToPDFArgs().
		SetLandscape(false).
		SetPrintBackground(true).
		SetMarginTop(0).
		SetMarginBottom(0).
		SetMarginLeft(0).
		SetMarginRight(0).
		SetPaperWidth(pdfWidth).
		SetPaperHeight(pdfHeight)

	print, _ := cli.Page.PrintToPDF(ctx, printToPDFArgs)
	pdfOutput := path.Join(htmllocaldir, "output.pdf")
	if err = ioutil.WriteFile(pdfOutput, print.Data, 0644); err != nil {
		return
	}

	fmt.Printf("Saved pdf: %s\n", pdfOutput)

	return nil
}
