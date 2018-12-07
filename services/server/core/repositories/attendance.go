package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// AttendanceRepository friends repository
type AttendanceRepository struct{}

// NewAttendanceRepository init UserProfile repository
func NewAttendanceRepository() *AttendanceRepository {
	return &AttendanceRepository{}
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
