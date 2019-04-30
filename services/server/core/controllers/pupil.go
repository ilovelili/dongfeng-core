package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PupilController pupil controller
type PupilController struct {
	repository      *repositories.PupilRepository
	classcontroller *ClassController
}

// NewPupilController new pupil controller
func NewPupilController() *PupilController {
	return &PupilController{
		repository:      repositories.NewPupilRepository(),
		classcontroller: NewClassController(),
	}
}

// GetPupils get pupils
func (c *PupilController) GetPupils(class, year string) ([]*models.Pupil, error) {
	return c.repository.Select(class, year)
}

// UpdatePupil update pupil
func (c *PupilController) UpdatePupil(pupil *models.Pupil) error {
	classes, err := c.classcontroller.GetClasses()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetClasses)
	}

	for _, class := range classes {
		if pupil.Class == class.Name {
			if err = c.repository.Update(pupil); err != nil {
				return utils.NewError(errorcode.CoreFailedToUpdatePupils)
			}
			return nil
		}
	}

	return utils.NewError(errorcode.CoreInvalidClass)
}

// UpdatePupils delete insert pupils
func (c *PupilController) UpdatePupils(pupils []*proto.Pupil) error {
	return c.repository.DeleteInsert(pupils)
}
