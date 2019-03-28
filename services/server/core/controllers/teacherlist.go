package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherlistController teacherlist controller
type TeacherlistController struct {
	repository *repositories.TeacherlistRepository
}

// NewTeacherlistController new controller
func NewTeacherlistController() *TeacherlistController {
	return &TeacherlistController{
		repository: repositories.NewTeacherlistRepository(),
	}
}

// GetTeacherlists get teacherlists
func (c *TeacherlistController) GetTeacherlists(class, year string) ([]*models.Teacherlist, error) {
	return c.repository.Select(class, year)
}

// UpdateTeacherlists update teacherlists
func (c *TeacherlistController) UpdateTeacherlists(items []*proto.TeacherlistItem) error {
	return c.repository.DeleteInsert(items)
}
