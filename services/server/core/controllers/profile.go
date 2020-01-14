package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// ProfileController profile controller
type ProfileController struct {
	repository      *repositories.ProfileRepository
	pupilcontroller *PupilController
}

// NewProfileController new controller
func NewProfileController() *ProfileController {
	return &ProfileController{
		repository:      repositories.NewProfileRepository(),
		pupilcontroller: NewPupilController(),
	}
}

// GetProfile get profile
func (c *ProfileController) GetProfile(year, class, name, date string) (profile *models.Profile, err error) {
	return c.repository.Select(year, class, name, date)
}

// GetPrevProfile get prev profile
func (c *ProfileController) GetPrevProfile(year, class, name, date string) (profile *models.Profile, err error) {
	return c.repository.SelectPrev(year, class, name, date)
}

// GetNextProfile get next profile
func (c *ProfileController) GetNextProfile(year, class, name, date string) (profile *models.Profile, err error) {
	return c.repository.SelectNext(year, class, name, date)
}

// GetProfiles get profiles
func (c *ProfileController) GetProfiles(year, class, name string) (profiles []*models.Profile, err error) {
	return c.repository.SelectAll(year, class, name)
}

// SaveProfile save profile
func (c *ProfileController) SaveProfile(profile *models.Profile) error {
	// save by class
	if profile.Name == profile.Class {
		profiles := []*models.Profile{profile}
		pupils, err := c.pupilcontroller.GetPupils(profile.Class, profile.Year)
		if err != nil {
			return err
		}

		for _, pupil := range pupils {
			profiles = append(profiles, &models.Profile{
				Year:      profile.Year,
				Class:     profile.Class,
				Name:      pupil.Name,
				Date:      profile.Date,
				Profile:   profile.Profile,
				Template:  profile.Template,
				CreatedBy: profile.CreatedBy,
				Enabled:   profile.Enabled,
			})
		}

		return c.repository.UpsertAll(profiles)
	}

	return c.repository.Upsert(profile)
}

// InsertProfile insert profile
func (c *ProfileController) InsertProfile(profile *models.Profile) error {
	// insert by class
	if profile.Name == profile.Class {
		profiles := []*models.Profile{profile}
		pupils, err := c.pupilcontroller.GetPupils(profile.Class, profile.Year)
		if err != nil {
			return err
		}

		for _, pupil := range pupils {
			profiles = append(profiles, &models.Profile{
				Year:      profile.Year,
				Class:     profile.Class,
				Name:      pupil.Name,
				Date:      profile.Date,
				Profile:   profile.Profile,
				Template:  profile.Template,
				CreatedBy: profile.CreatedBy,
				Enabled:   profile.Enabled,
			})
		}

		return c.repository.InsertAll(profiles)
	}

	return c.repository.Insert(profile)
}

// DeleteProfile delete profile
func (c *ProfileController) DeleteProfile(profile *models.Profile) error {
	// delete by class
	if profile.Name == profile.Class {
		profiles := []*models.Profile{profile}
		pupils, err := c.pupilcontroller.GetPupils(profile.Class, profile.Year)
		if err != nil {
			return err
		}

		for _, pupil := range pupils {
			profiles = append(profiles, &models.Profile{
				Year:      profile.Year,
				Class:     profile.Class,
				Name:      pupil.Name,
				Date:      profile.Date,
				Profile:   profile.Profile,
				Template:  profile.Template,
				CreatedBy: profile.CreatedBy,
				Enabled:   profile.Enabled,
			})
		}

		return c.repository.DeleteAll(profiles)
	}

	return c.repository.Delete(profile)
}
