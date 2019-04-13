package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// AttendanceController attendance controller
type AttendanceController struct {
	repository *repositories.AttendanceRepository
}

// NewAttendanceController new controller
func NewAttendanceController() *AttendanceController {
	return &AttendanceController{
		repository: repositories.NewAttendanceRepository(),
	}
}

// SelectAttendances select attendence
func (c *AttendanceController) SelectAttendances(year, from, to, class, name string) (attendances []*models.Attendance, err error) {
	return c.repository.Select(year, from, to, class, name)
}

// UpdateAttendances update attendence
func (c *AttendanceController) UpdateAttendances(attendances []*models.Attendance) error {
	return c.repository.DeleteInsert(attendances)
}
