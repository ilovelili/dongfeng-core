package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// NamelistRepository name list repository
type NamelistRepository struct{}

// NewNamelistRepository init name list repository
func NewNamelistRepository() *NamelistRepository {
	return &NamelistRepository{}
}

// Select select namelist
func (r *NamelistRepository) Select(class, year string) (namelists []*models.Namelist, err error) {
	querybuilder := Table("namelists").Alias("n").Where()
	var query string

	if class == "" && year == "" {
		querybuilder = querybuilder.Eq("1", "1")
	} else {
		if class != "" {
			querybuilder = querybuilder.Eq("n.class", class)
		}

		if year != "" {
			querybuilder = querybuilder.Eq("n.year", year)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&namelists)
	return
}

// DeleteInsert delete insert namelist
func (r *NamelistRepository) DeleteInsert(namelists []*proto.NamelistItem) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for idx, namelist := range namelists {
		year, class, names, createdBy := namelist.GetYear(), namelist.GetClass(), namelist.GetNames(), namelist.GetCreatedBy()
		if idx == 0 {
			_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteNamelist('%s','%s')", year, class))
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		for _, name := range names {
			err = session().InsertTx(tx, &models.Namelist{
				Year:      year,
				Class:     class,
				Name:      name.Name,
				CreatedBy: createdBy,
			})

			if err != nil {
				session().Rollback(tx)
				return
			}
		}
	}

	return session().Commit(tx)
}
