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

// Insert insert attendances
func (r *AttendanceRepository) Insert(attendances []*models.Attendance) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// insert by loop (use bulk insert?)
	for _, attendance := range attendances {
		err = session().InsertTx(tx, attendance)
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
