package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ClassController class controller
type ClassController struct {
	repository *repositories.ClassRepository
}

// NewClassController new class controller
func NewClassController() *ClassController {
	return &ClassController{
		repository: repositories.NewClassRepository(),
	}
}

// GetClasses get Classes
func (c *ClassController) GetClasses() ([]*models.Class, error) {
	return c.repository.Select()
}

// UpdateClasses update Classes
func (c *ClassController) UpdateClasses(classes []*proto.Class) error {
	return c.repository.DeleteInsert(classes)
}
