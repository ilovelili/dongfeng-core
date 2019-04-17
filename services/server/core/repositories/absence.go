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
func (r *AbsenceRepository) Select(year, from, to, class, name string) (absences []*models.Absence, err error) {
	var query string
	table := Table("absences").Alias("a")
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
	if err = session().Find(query, nil).All(&absences); err != nil && norows(err) {
		err = nil
	}

	return
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
