package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// ProfileTemplateController growth profile template controller
type ProfileTemplateController struct {
	repository *repositories.ProfileTemplateRepository
}

// NewProfileTemplateController new growth profile template controller
func NewProfileTemplateController() *ProfileTemplateController {
	return &ProfileTemplateController{
		repository: repositories.NewProfileTemplateRepository(),
	}
}

// GetProfileTemplates get profile templatees
func (c *ProfileTemplateController) GetProfileTemplates() ([]*models.ProfileTemplate, error) {
	return c.repository.Select()
}

// UpdateProfileTemplates update profile templatees
func (c *ProfileTemplateController) UpdateProfileTemplates(template *models.ProfileTemplate) error {
	return c.repository.Upsert(template)
}
