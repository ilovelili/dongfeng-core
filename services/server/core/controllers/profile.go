package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// ProfileController profile controller
type ProfileController struct {
	repository *repositories.ProfileRepository
}

// NewProfileController new controller
func NewProfileController() *ProfileController {
	return &ProfileController{
		repository: repositories.NewProfileRepository(),
	}
}

// GetProfile get profile
func (c *ProfileController) GetProfile(year, class, name string) (profile *models.Profile, err error) {
	return c.repository.Select(year, class, name)
}

// SaveProfile save profile
func (c *ProfileController) SaveProfile(profile *models.Profile) error {
	return c.repository.Upsert(profile)
}
