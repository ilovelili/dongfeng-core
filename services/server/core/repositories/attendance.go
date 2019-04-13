package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// AttendanceRepository friends repository
type AttendanceRepository struct{}

// NewAttendanceRepository init UserProfile repository
func NewAttendanceRepository() *AttendanceRepository {
	return &AttendanceRepository{}
}

// Select select attendances
func (r *AttendanceRepository) Select(year, from, to, class, name string) (attendances []*models.Attendance, err error) {
	var query string
	table := Table("attendances").Alias("a")
	if from == "" && to == "" && class == "" && name == "" {
		query = table.Sql()
	} else {
		querybuilder := table.Where()

		if from != "" && to != "" && from > to {
			err = fmt.Errorf("invalid parameter")
			return
		}

		if year != "" {
			querybuilder.Eq("a.year", year)
		}
		if from != "" {
			querybuilder.Gte("a.date", from)
		}
		if to != "" {
			querybuilder.Lte("a.date", to)
		}
		if class != "" {
			querybuilder.Eq("a.class", class)
		}
		if name != "" {
			querybuilder.Eq("a.name", name)
		}

		query = querybuilder.Sql()
	}

	// no rows is actually not an error
	if err = session().Find(query, nil).All(&attendances); err != nil && norows(err) {
		err = nil
	}

	return
}

// DeleteInsert deleteinsert attendances
func (r *AttendanceRepository) DeleteInsert(attendances []*models.Attendance) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// upsert by loop
	for idx, attendance := range attendances {
		year, class, date := attendance.Year, attendance.Class, attendance.Date
		if idx == 0 {
			_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteAttendances('%s','%s','%s')", year, class, date))
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		err = session().InsertTx(tx, attendance)
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
