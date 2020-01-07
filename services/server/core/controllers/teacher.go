package controllers

import (
	"strings"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherController teacher controller
type TeacherController struct {
	repository      *repositories.TeacherRepository
	classcontroller *ClassController
	rolecontroller  *RoleController
}

// NewTeacherController new teacher controller
func NewTeacherController() *TeacherController {
	return &TeacherController{
		repository:      repositories.NewTeacherRepository(),
		classcontroller: NewClassController(),
		rolecontroller:  NewRoleController(),
	}
}

// GetTeachers get teacherlists
func (c *TeacherController) GetTeachers(class, year string) ([]*models.Teacher, error) {
	return c.repository.Select(class, year)
}

// UpdateTeacher update teacher
func (c *TeacherController) UpdateTeacher(teacher *models.Teacher) error {
	// empty class is allowed
	if teacher.Class == "" {
		if err := c.repository.Update(teacher); err != nil {
			return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
		}
		return nil
	}

	classsegments := strings.Split(teacher.Class, "|")
	if len(classsegments) == 0 {
		return utils.NewError(errorcode.CoreInvalidClass)
	}

	classes, err := c.classcontroller.GetClasses()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetClasses)
	}

	foundcount := 0
	for _, classsegment := range classsegments {
		for _, class := range classes {
			if classsegment == class.Name {
				foundcount++
				break
			}
		}
	}

	// all class found in database, valid
	if foundcount == len(classsegments) {
		if err := c.repository.Update(teacher); err != nil {
			return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
		}

		// update role
		if teacher.Email != "" {
			var r string
			if teacher.Role == nil {
				r = "教师"
			} else {
				r = *teacher.Role
			}
			if err := c.rolecontroller.SaveRole(&models.Role{
				User: teacher.Email,
				Role: r,
			}); err != nil {
				return utils.NewError(errorcode.CoreFailedToUpdateRole)
			}
		}

		return nil
	}

	return utils.NewError(errorcode.CoreInvalidClass)
}

// UpdateTeachers update teachers
func (c *TeacherController) UpdateTeachers(teachers []*proto.Teacher) error {
	if err := c.repository.DeleteInsert(teachers); err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
	}

	// update role
	for _, teacher := range teachers {
		if teacher.GetEmail() != "" {
			r := teacher.GetRole()
			if r == "" {
				r = "教师"
			}

			if err := c.rolecontroller.repository.Upsert(&models.Role{
				User: teacher.GetEmail(),
				Role: r,
			}); err != nil {
				return utils.NewError(errorcode.CoreFailedToUpdateRole)
			}
		}
	}

	return nil
}
