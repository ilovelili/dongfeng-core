package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// HolidayRepository holiday repository
type HolidayRepository struct{}

// NewHolidayRepository init holiday repository
func NewHolidayRepository() *HolidayRepository {
	return &HolidayRepository{}
}

// Select select holidays
func (r *HolidayRepository) Select() (holidays []*models.Holiday, err error) {
	query := Table("holidays").Sql()
	err = session().Find(query, nil).All(&holidays)
	return
}
