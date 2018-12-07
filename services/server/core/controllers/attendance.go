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

// Save save attendence
func (c *AttendanceController) Save(attendances []*models.Attendance) error {
	return c.repository.Upsert(attendances)
}
