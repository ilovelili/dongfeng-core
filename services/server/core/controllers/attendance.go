package controllers

import (
	"time"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// AttendanceController attendance controller
type AttendanceController struct {
	repository        *repositories.AbsenceRepository
	holidaycontroller *HolidayController
	pupilcontroller   *PupilController
}

// NewAttendanceController new controller
func NewAttendanceController() *AttendanceController {
	return &AttendanceController{
		repository:        repositories.NewAbsenceRepository(),
		holidaycontroller: NewHolidayController(),
		pupilcontroller:   NewPupilController(),
	}
}

// SelectAttendances select attendence
func (c *AttendanceController) SelectAttendances(year, from, to, class, name string) (attendances *models.Attendances, err error) {
	// get one month as default
	if from == "" {
		from = time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	}

	if to == "" {
		to = time.Now().Format("2006-01-02")
	}

	start, err := time.Parse("2006-01-02", from)
	if err != nil {
		return
	}

	end, err := time.Parse("2006-01-02", to)
	if err != nil {
		return
	}

	_attendances := []*models.Attendance{}
	holidaytypes := []*models.HolidayType{}

	absences, err := c.repository.Select(year, from, to, class, name)
	if err != nil {
		return
	}

	holidays, err := c.holidaycontroller.GetHolidaysInString(from, to)
	if err != nil {
		return
	}

	pupils, err := c.pupilcontroller.GetPupils(class, year)
	if err != nil {
		return
	}

	for d := start; d.Before(end.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		date := d.Format("2006-01-02")
		// skip if it's holiday
		if sharedlib.ContainString(holidays, date) {
			holidaytypes = append(holidaytypes, &models.HolidayType{
				Date: date,
				Type: 2,
			})

			continue
		}

		// skip if it's weekend
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			holidaytypes = append(holidaytypes, &models.HolidayType{
				Date: date,
				Type: 1,
			})

			continue
		}

		for _, pupil := range pupils {
			pupilinabsence := false
			for _, absence := range absences {
				// if pupil name not in absence list, put it in attendance
				if absence.Date == date && pupil.Name == absence.Name && pupil.Year == absence.Year && pupil.Class == absence.Class {
					pupilinabsence = true
					break
				}
			}

			if !pupilinabsence {
				_attendances = append(_attendances, &models.Attendance{
					Year:  pupil.Year,
					Date:  date,
					Class: pupil.Class,
					Name:  pupil.Name,
				})
			}
		}
	}

	attendances = &models.Attendances{
		Attendances: _attendances,
		Holidays:    holidaytypes,
	}

	return attendances, nil
}

// UpdateAbsences update absences
func (c *AttendanceController) UpdateAbsences(absences []*models.Absence) error {
	return c.repository.DeleteInsert(absences)
}
