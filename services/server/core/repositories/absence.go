package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// AbsenceRepository absence repository
type AbsenceRepository struct{}

// NewAbsenceRepository init absence repository
func NewAbsenceRepository() *AbsenceRepository {
	return &AbsenceRepository{}
}

// Select select absences
func (r *AbsenceRepository) Select(year, from, to, class string) (absences []*models.Absence, err error) {
	var query string
	table := Table("absences").Alias("a")
	if from == "" && to == "" && class == "" {
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

		query = querybuilder.Sql()
	}

	// no rows is actually not an error
	if err = session().Find(query, nil).All(&absences); err != nil && norows(err) {
		err = nil
	}

	return
}

// Update deleteinsert attendances
func (r *AbsenceRepository) Update(absences []*models.Absence, attendances []*models.Absence) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for _, attendance := range attendances {
		year, class, date, name := attendance.Year, attendance.Class, attendance.Date, attendance.Name
		// delete from absence table
		_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteAbsence('%s','%s','%s', '%s')", year, class, date, name))
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	for _, absence := range absences {
		query := Table("absences").Alias("a").Where().
			Eq("a.year", absence.Year).
			Eq("a.class", absence.Class).
			Eq("a.date", absence.Date).
			Eq("a.name", absence.Name).
			Sql()

		var a models.Absence
		err := session().Find(query, nil).Single(&a)
		if err != nil || 0 == a.ID {
			err = session().InsertTx(tx, absence)
		} else {
			absence.ID = a.ID
			err = session().UpdateTx(tx, absence)
		}

		if err != nil {
			session().Rollback(tx)
			return err
		}
	}

	return session().Commit(tx)
}

// DeleteInsert deleteinsert attendances
func (r *AbsenceRepository) DeleteInsert(absences []*models.Absence) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// upsert by loop
	for idx, absence := range absences {
		year, class, date := absence.Year, absence.Class, absence.Date
		if idx == 0 {
			_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteAbsences('%s','%s','%s')", year, class, date))
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		err = session().InsertTx(tx, absence)
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
