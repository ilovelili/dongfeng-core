package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherController teacher controller
type TeacherController struct {
	repository *repositories.TeacherRepository
}

// NewTeacherController new teacher controller
func NewTeacherController() *TeacherController {
	return &TeacherController{
		repository: repositories.NewTeacherRepository(),
	}
}

// GetTeachers get teacherlists
func (c *TeacherController) GetTeachers(class, year string) ([]*models.Teacher, error) {
	return c.repository.Select(class, year)
}

// UpdateTeacher update teacher
func (c *TeacherController) UpdateTeacher(teacher *models.Teacher) error {
	return c.repository.Update(teacher)
}

// UpdateTeachers update teachers
func (c *TeacherController) UpdateTeachers(teachers []*proto.Teacher) error {
	return c.repository.DeleteInsert(teachers)
}
