package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PupilRepository pupil repository
type PupilRepository struct{}

// NewPupilRepository init pupil repository
func NewPupilRepository() *PupilRepository {
	return &PupilRepository{}
}

// Select select pupils
func (r *PupilRepository) Select(class, year string) (pupils []*models.Pupil, err error) {
	querybuilder := Table("pupils").Alias("p").Where()
	var query string

	if class == "" && year == "" {
		querybuilder = querybuilder.Eq("1", "1")
	} else {
		if class != "" {
			querybuilder = querybuilder.Eq("p.class", class)
		}

		if year != "" {
			querybuilder = querybuilder.Eq("p.year", year)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&pupils)
	return
}

// DeleteInsert delete insert pupils
func (r *PupilRepository) DeleteInsert(pupils []*proto.Pupil) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for idx, pupil := range pupils {
		year, class, name, createdBy := pupil.GetYear(), pupil.GetClass(), pupil.GetName(), pupil.GetCreatedBy()
		if idx == 0 {
			_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeletePupils('%s','%s')", year, class))
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		err = session().InsertTx(tx, &models.Pupil{
			Year:      year,
			Class:     class,
			Name:      name,
			CreatedBy: createdBy,
		})

		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
