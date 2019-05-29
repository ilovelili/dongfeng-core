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
func (c *ProfileController) GetProfile(year, class, name, date string) (profile *models.Profile, err error) {
	return c.repository.Select(year, class, name, date)
}

// GetProfiles get profiles
func (c *ProfileController) GetProfiles(year, class, name string) (profiles []*models.Profile, err error) {
	return c.repository.SelectAll(year, class, name)
}

// GetProfileNames get profile names
func (c *ProfileController) GetProfileNames(year, class, name, date string) (names []string, err error) {
	profiles, err := c.repository.SelectNames(year, class, name, date)
	if err != nil {
		return
	}

	for _, profile := range profiles {
		names = append(names, profile.Name)
	}
	return
}

// GetProfileDates get profile dates
func (c *ProfileController) GetProfileDates(year, class, name, date string) (dates []string, err error) {
	profiles, err := c.repository.SelectDates(year, class, name, date)
	if err != nil {
		return
	}

	for _, profile := range profiles {
		dates = append(dates, profile.Date)
	}
	return
}

// SaveProfile save profile
func (c *ProfileController) SaveProfile(profile *models.Profile) error {
	return c.repository.Upsert(profile)
}
