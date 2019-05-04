package repositories

import (
	"fmt"
	"strings"

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

// Update update pupil
func (r *PupilRepository) Update(pupil *models.Pupil) (err error) {
	return session().Update(pupil)
}

// DeleteInsert delete insert pupils
func (r *PupilRepository) DeleteInsert(pupils []*proto.Pupil) (err error) {
	pupilsmap := make(map[string][]*proto.Pupil)
	for _, pupil := range pupils {
		key := fmt.Sprintf("%s_%s", pupil.GetYear(), pupil.GetClass())
		if pupils, ok := pupilsmap[key]; !ok {
			pupilsmap[key] = []*proto.Pupil{pupil}
		} else {
			pupilsmap[key] = append(pupils, pupil)
		}
	}

	tx, err := session().Begin()
	if err != nil {
		return
	}

	for k := range pupilsmap {
		segments := strings.Split(k, "_")
		if len(segments) != 2 {
			err = fmt.Errorf("invalid key")
			return
		}

		year, class := segments[0], segments[1]
		_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeletePupils('%s','%s')", year, class))
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	for _, pupil := range pupils {
		year, class, name, createdBy := pupil.GetYear(), pupil.GetClass(), pupil.GetName(), pupil.GetCreatedBy()
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
