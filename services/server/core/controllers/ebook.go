package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	oss "github.com/ilovelili/aliyun-client/oss"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

var (
	config = utils.GetConfig()
)

// EbookController ebook controller
type EbookController struct {
	repository *repositories.EbookRepository
	svc        *oss.Service
}

// NewEbookController new controller
func NewEbookController() *EbookController {
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

// GetEbooks get ebooks
func (c *EbookController) GetEbooks(year, class, name string) ([]*models.Ebook, error) {
	ebooks, err := c.repository.Select(year, class, name)
	if err != nil {
		return nil, err
	}

	ebookmap := make(map[string][]string)
	for _, ebook := range ebooks {
		key := fmt.Sprintf("%s_%s_%s", ebook.Year, ebook.Class, ebook.Name)
		if dates, ok := ebookmap[key]; ok {
			ebookmap[key] = append(dates, ebook.Date)
		} else {
			ebookmap[key] = []string{ebook.Date}
		}
	}

	results := []*models.Ebook{}
EbookLoop:
	for _, ebook := range ebooks {
		for _, result := range results {
			if result.Year == ebook.Year && result.Class == ebook.Class && result.Name == ebook.Name {
				continue EbookLoop
			}
		}

		for k, v := range ebookmap {
			segments := strings.Split(k, "_")
			if len(segments) != 3 {
				return nil, fmt.Errorf("invalid key")
			}

			if ebook.Year == segments[0] && ebook.Class == segments[1] && ebook.Name == segments[2] {
				ebook.Dates = v
				results = append(results, ebook)
			}
		}
	}

	return results, nil
}

// SaveEbook save ebook
func (c *EbookController) SaveEbook(ebook *models.Ebook) error {
	ebook.ResolveHash()
	ebook.Converted = false
	dirty, err := c.repository.Upsert(ebook, false)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveEbook)
	}

	// if dirty
	if dirty {
		// upload to storage
		if err = c.uploadToStorage(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToUploadEbookToCloud)
		}

		// convert to pdf
		if err = c.convert(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToConvertEbookHTML)
		}

		// merge to ebook
		if err = c.merge(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToMergeEbook)
		}

		// set converted to true if everything goes smoothly
		ebook.Converted = true
		_, err = c.repository.Upsert(ebook, true)
		if err != nil {
			return utils.NewError(errorcode.CoreFailedToSaveEbook)
		}
	}

	return nil
}

// uploadToCloudStorage upload css folder and index.html to local (or aliyun oss later)
func (c *EbookController) uploadToStorage(ebook *models.Ebook) error {
	htmllocaldir := path.Join(config.Ebook.OriginDir, ebook.Year, ebook.Class, ebook.Name, ebook.Date)
	csslocaldir := path.Join(htmllocaldir, "css")

	_, err := os.Stat(csslocaldir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(csslocaldir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	csslocalfile := path.Join(csslocaldir, "style.css")
	err = ioutil.WriteFile(csslocalfile, []byte(ebook.ResolveCloudCSS()), os.ModePerm)
	if err != nil {
		return err
	}

	htmllocalfile := path.Join(htmllocaldir, "index.html")
	err = ioutil.WriteFile(htmllocalfile, []byte(ebook.ResolveCloudHTML()), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// convert ebook to pdf and jpg using chrome headless
func (c *EbookController) convert(ebook *models.Ebook) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	chromeDevTool := os.Getenv("CHROME_DEV_TOOL")
	if chromeDevTool == "" {
		chromeDevTool = "http://127.0.0.1:9222"
	}

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

	// Enable the runtime
	if err = cli.Runtime.Enable(ctx); err != nil {
		return
	}

	// Enable the network
	if err = cli.Network.Enable(ctx, network.NewEnableArgs()); err != nil {
		return
	}

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = cli.Page.Enable(ctx); err != nil {
		return
	}

	htmllocaldir := path.Join(config.Ebook.OriginDir, ebook.Year, ebook.Class, ebook.Name, ebook.Date)
	// Create the Navigate arguments
	navArgs := page.NewNavigateArgs(fmt.Sprintf("file://%s", path.Join(htmllocaldir, "index.html")))
	nav, err := cli.Page.Navigate(ctx, navArgs)
	if err != nil {
		return
	}

	// wait till image loaded
	time.Sleep(time.Duration(config.Ebook.ImageLoadTimeout) * time.Second)

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
		SetPaperWidth(config.Ebook.Width).
		SetPaperHeight(config.Ebook.Height)

	print, _ := cli.Page.PrintToPDF(ctx, printToPDFArgs)
	pdfOutput := path.Join(htmllocaldir, "output.pdf")
	if err = ioutil.WriteFile(pdfOutput, print.Data, 0644); err != nil {
		return
	}

	fmt.Printf("Saved pdf: %s\n", pdfOutput)

	// move to dest dir
	pdfdestdir := path.Join(config.Ebook.PDFDestDir, ebook.Year, ebook.Class, ebook.Name)
	_, err = os.Stat(pdfdestdir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(pdfdestdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err = os.Rename(pdfOutput, path.Join(pdfdestdir, fmt.Sprintf("%s.pdf", ebook.Date))); err != nil {
		return err
	}

	imgdestdir := path.Join(config.Ebook.ImageDestDir, ebook.Year, ebook.Class, ebook.Name)
	_, err = os.Stat(imgdestdir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(imgdestdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = os.Rename(imgOutput, path.Join(imgdestdir, fmt.Sprintf("%s.jpg", ebook.Date)))
	return err

}

// merge merge pdf files into ebook
func (c *EbookController) merge(ebook *models.Ebook) (err error) {
	// check if pdftk installed or not
	_, err = exec.LookPath("pdftk")
	if err != nil {
		return
	}

	filepathmap := make(map[string][]string)
	targetdir := config.Ebook.MergeTargetDir
	destdir := path.Join(config.Ebook.MergeDestDir, ebook.Class, ebook.Name)

	err = filepath.Walk(targetdir, func(filepath string, info os.FileInfo, err error) error {
		// target
		if !info.IsDir() && path.Ext(info.Name()) == ".pdf" {
			key := path.Dir(filepath)
			// select the target file with corresponding class and name
			if strings.Index(key, fmt.Sprintf("/%s/%s", ebook.Class, ebook.Name)) == -1 {
				return nil
			}

			// ignore the dest file
			if strings.Index(key, config.Ebook.MergeDestDir) > -1 {
				return nil
			}

			if paths, ok := filepathmap[key]; ok {
				filepathmap[key] = append(paths, filepath)
			} else {
				filepathmap[key] = []string{filepath}
			}
		}
		return nil
	})

	if err != nil {
		return
	}

	// first clear dest dir
	os.RemoveAll(destdir)
	_, err = os.Stat(destdir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(destdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for dir, filepaths := range filepathmap {
		// sort pdf by date
		sort.Strings(filepaths)
		// https://stackoverflow.com/questions/31467153/golang-failed-exec-command-that-works-in-terminal
		// cmdline := fmt.Sprintf("pdftk %s cat output merge.pdf", path.Join(filepath, "*.pdf"))
		pdffiles := strings.Join(filepaths, " ")
		cmdline := fmt.Sprintf("pdftk %s cat output %s", pdffiles, path.Join(dir, "merge.pdf"))
		args := strings.Split(cmdline, " ")
		cmd := exec.Command(args[0], args[1:]...)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}

		// move to dest
		// 电子书_${this.currentName}_${this.currentClass}_${this.currentYear}学年.pdf
		err = os.Rename(path.Join(dir, "merge.pdf"), path.Join(destdir, fmt.Sprintf("电子书_%s_%s_%s学年.pdf", ebook.Name, ebook.Class, ebook.Year)))
		if err != nil {
			return
		}
	}

	// loop dest dir and merge again to generate the full year ebook
	destfilepathmap := make(map[string][]string)
	err = filepath.Walk(config.Ebook.MergeDestDir, func(filepath string, info os.FileInfo, err error) error {
		if !info.IsDir() && path.Ext(info.Name()) == ".pdf" {
			key := path.Dir(filepath)
			if paths, ok := destfilepathmap[key]; ok {
				destfilepathmap[key] = append(paths, filepath)
			} else {
				destfilepathmap[key] = []string{filepath}
			}
		}
		return nil
	})

	if err != nil {
		return
	}

	for dir, filepaths := range destfilepathmap {
		sort.Strings(filepaths)

		pdffiles := strings.Join(filepaths, " ")
		// move to dest
		segments := strings.Split(dir, "/")
		class, name := segments[len(segments)-2], segments[len(segments)-1]
		// 电子书_${this.currentName}_${this.currentClass}_全期间.pdf
		cmdline := fmt.Sprintf("pdftk %s cat output %s", pdffiles, path.Join(dir, fmt.Sprintf("电子书_%s_%s_全期间.pdf", name, class)))
		args := strings.Split(cmdline, " ")
		cmd := exec.Command(args[0], args[1:]...)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
	}

	return
}
