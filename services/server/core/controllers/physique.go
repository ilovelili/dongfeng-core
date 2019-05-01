package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PhysiqueController physique controller
type PhysiqueController struct {
	repository      *repositories.PhysiqueRepository
	pupilcontroller *PupilController
}

// NewPhysiqueController new physique controller
func NewPhysiqueController() *PhysiqueController {
	return &PhysiqueController{
		repository:      repositories.NewPhysiqueRepository(),
		pupilcontroller: NewPupilController(),
	}
}

// GetPhysiques get physiques
func (c *PhysiqueController) GetPhysiques(class, year, name string) ([]*models.Physique, error) {
	return c.repository.Select(class, year, name)
}

// UpdatePhysique update physique
func (c *PhysiqueController) UpdatePhysique(physique *models.Physique) error {
	pupils, err := c.pupilcontroller.GetPupils(physique.Class, physique.Year)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiques)
	}

	for _, pupil := range pupils {
		if physique.Name == pupil.Name {
			//
			if err = c.repository.Update(physique); err != nil {
				return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
			}
			return nil
		}
	}

	return utils.NewError(errorcode.CoreInvalidPupil)
}

// UpdatePhysiques delete insert physique
func (c *PhysiqueController) UpdatePhysiques(physique []*proto.Physique) error {
	return c.repository.DeleteInsert(physique)
}

func (c *PhysiqueController) ResolvePhysique(physique *models.Physique) error {
	p.resolveAge()
	p.resolveBMI()
}
