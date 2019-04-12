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

// Upsert upsert attendances
func (r *AttendanceRepository) Upsert(attendances []*models.Attendance) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// upsert by loop
	for _, attendance := range attendances {
		query := Table("attendances").Alias("a").Project("a.id").Where().
			Eq("a.year", attendance.Year).
			Eq("a.date", attendance.Date).
			Eq("a.class", attendance.Class).
			Eq("a.name", attendance.Name).
			Sql()

		var id int64
		err := session().Find(query, nil).Scalar(&id)
		if err != nil || 0 == id {
			err = session().InsertTx(tx, attendance)
		} else {
			attendance.ID = id
			err = session().UpdateTx(tx, attendance)
		}

		if err != nil {
			session().Rollback(tx)
			return err
		}
	}

	return session().Commit(tx)
}
