package controllers

import (
	"time"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// HolidayController pupil controller
type HolidayController struct {
	repository *repositories.HolidayRepository
}

// NewHolidayController new holiday controller
func NewHolidayController() *HolidayController {
	return &HolidayController{
		repository: repositories.NewHolidayRepository(),
	}
}

// GetHolidays get pupils
func (c *HolidayController) GetHolidays() ([]*models.Holiday, error) {
	return c.repository.Select()
}

// GetHolidaysInString get holidays in string format
func (c *HolidayController) GetHolidaysInString(from, to string) ([]string, error) {
	datefrom, err := time.Parse("2006-01-02", from)
	if err != nil {
		return []string{}, err
	}

	dateto, err := time.Parse("2006-01-02", to)
	if err != nil {
		return []string{}, err
	}

	_holidays, err := c.GetHolidays()
	if err != nil {
		return []string{}, err
	}

	result := []string{}
	for _, h := range _holidays {
		start, err := time.Parse("2006-01-02", h.From)
		if err != nil {
			return []string{}, err
		}

		end, err := time.Parse("2006-01-02", h.To)
		if err != nil {
			return []string{}, err
		}

		if start.After(dateto) || end.Before(datefrom) {
			continue
		}

		if start.Before(datefrom) {
			start = datefrom
		}

		if end.After(dateto) {
			end = dateto
		}

		for d := start; d.Before(end.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
			result = append(result, d.Format("2006-01-02"))
		}
	}

	return result, nil
}
