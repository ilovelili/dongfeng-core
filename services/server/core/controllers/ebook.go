package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	oss "github.com/ilovelili/aliyun-client/oss"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
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
	dirty, err := c.repository.Upsert(ebook)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveEbook)
	}

	// if dirty
	if dirty {
		// upload to storage
		if err = c.uploadToStorage(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToUploadEbookToCloud)
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
