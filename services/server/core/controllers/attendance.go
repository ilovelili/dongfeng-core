package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
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

	absences, err := c.repository.Select(year, from, to, class)
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

			if name == "" || name == pupil.Name {
				_attendances = append(_attendances, &models.Attendance{
					Year:           pupil.Year,
					Date:           date,
					Class:          pupil.Class,
					Name:           pupil.Name,
					AttendanceFlag: !pupilinabsence,
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

// UpdateAbsence update single absence
func (c *AttendanceController) UpdateAbsence(absences []*models.Absence, attendances []*models.Absence) error {
	return c.repository.Update(absences, attendances)
}

// UpdateAbsences update absences
func (c *AttendanceController) UpdateAbsences(absences []*models.Absence) error {
	return c.repository.DeleteInsert(absences)
}

// CountAttendance count attendance
func (c *AttendanceController) CountAttendance(from, to, class string) (attendancecount *models.AttendanceCount, err error) {
	if from != "" && to != "" && to < from {
		err = utils.NewError(errorcode.CoreInvalidAttendanceCountRequest)
		return
	}

	attendances, err := c.SelectAttendances("", from, to, class, "")
	if err != nil {
		err = utils.NewError(errorcode.CoreFailedToGetAttendances)
		return
	}

	attendancemap := make(map[string] /*date_class*/ int64 /*count*/)
	for _, attendance := range attendances.Attendances {
		// attended
		if attendance.AttendanceFlag {
			key := fmt.Sprintf("%s_%s", attendance.Date, attendance.Class)
			if v, ok := attendancemap[key]; ok {
				attendancemap[key] = v + 1
			} else {
				attendancemap[key] = 1
			}
		}
	}

	var sum, juniorsum, middlesum, seniorsum int64
	counts := []*models.ClassAttendanceCount{}

	for k, v := range attendancemap {
		segments := strings.Split(k, "_")
		if len(segments) != 2 {
			err = utils.NewError(errorcode.CoreFailedToGetAttendanceCount)
			return
		}
		date, class := segments[0], segments[1]
		counts = append(counts, &models.ClassAttendanceCount{
			Date:  date,
			Class: class,
			Count: v,
		})

		if c.isJuniorClass(class) {
			juniorsum += v
		} else if c.isMiddleClass(class) {
			middlesum += v
		} else if c.isSeniorClass(class) {
			seniorsum += v
		}

		sum += v
	}

	attendancecount = &models.AttendanceCount{
		ClassAttendanceCounts: counts,
		SeniorSum:             seniorsum,
		MiddleSum:             middlesum,
		JuniorSum:             juniorsum,
		Sum:                   sum,
	}

	return
}

func (c *AttendanceController) isJuniorClass(classname string) bool {
	return strings.Index(classname, "小") == 0 // must start with "小"
}

func (c *AttendanceController) isMiddleClass(classname string) bool {
	return strings.Index(classname, "中") == 0 // must start with "中"
}

func (c *AttendanceController) isSeniorClass(classname string) bool {
	return strings.Index(classname, "大") == 0 // must start with "大"
}
