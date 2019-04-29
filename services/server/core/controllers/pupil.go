package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PupilController pupil controller
type PupilController struct {
	repository *repositories.PupilRepository
}

// NewPupilController new pupil controller
func NewPupilController() *PupilController {
	return &PupilController{
		repository: repositories.NewPupilRepository(),
	}
}

// GetPupils get pupils
func (c *PupilController) GetPupils(class, year string) ([]*models.Pupil, error) {
	return c.repository.Select(class, year)
}

// UpdatePupil update pupil
func (c *PupilController) UpdatePupil(pupil *models.Pupil) error {
	return c.repository.Update(pupil)
}

// UpdatePupils delete insert pupils
func (c *PupilController) UpdatePupils(pupils []*proto.Pupil) error {
	return c.repository.DeleteInsert(pupils)
}
