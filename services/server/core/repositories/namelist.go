package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// NamelistRepository friends repository
type NamelistRepository struct{}

// NewNamelistRepository init UserProfile repository
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
			querybuilder = querybuilder.Eq("n.class_id", class)
		}

		if year != "" {
			querybuilder = querybuilder.Eq("n.year", year)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&namelists)
	return
}
