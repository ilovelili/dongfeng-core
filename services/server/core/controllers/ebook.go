package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
)

// EbookController ebook controller
type EbookController struct {
	repository *repositories.EbookRepository
}

// NewEbookController new controller
func NewEbookController() *EbookController {
	return &EbookController{
		repository: repositories.NewEbookRepository(),
	}
}

// SaveEbook save ebook
func (c *EbookController) SaveEbook(ebook *models.Ebook) error {
	ebook.ResolveHash()
	dirty, err := c.repository.Upsert(ebook)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveEbook)
	}

	// if dirty, upload to storage
	if dirty {
		if err = c.uploadToCloudStorage(ebook); err != nil {
			return utils.NewError(errorcode.CoreFailedToUploadEbookToCloud)
		}
	}

	return nil
}

// uploadToCloudStorage upload to cloud storage
func (c *EbookController) uploadToCloudStorage(ebook *models.Ebook) error {
	return nil
}
