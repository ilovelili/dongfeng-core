package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherlistRepository friends repository
type TeacherlistRepository struct{}

// NewTeacherlistRepository init UserProfile repository
func NewTeacherlistRepository() *TeacherlistRepository {
	return &TeacherlistRepository{}
}

// Select select Teacherlist
func (r *TeacherlistRepository) Select(class, year string) (teacherlists []*models.Teacherlist, err error) {
	querybuilder := Table("teacherlists").Alias("t").Where()
	var query string

	if class == "" && year == "" {
		querybuilder = querybuilder.Eq("1", "1")
	} else {
		if class != "" {
			querybuilder = querybuilder.Eq("t.class", class)
		}

		if year != "" {
			querybuilder = querybuilder.Eq("t.year", year)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&teacherlists)
	return
}

// DeleteInsert delete insert Teacherlist
func (r *TeacherlistRepository) DeleteInsert(teacherlists []*proto.TeacherlistItem) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for idx, teacherlist := range teacherlists {
		year, teachers, createdBy := teacherlist.GetYear(), teacherlist.GetItems(), teacherlist.GetCreatedBy()
		if idx == 0 {
			_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteTeacherlist('%s')", year))
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		for _, teacher := range teachers {
			for _, class := range teacher.Class {
				err = session().InsertTx(tx, &models.Teacherlist{
					Year:      year,
					Class:     class,
					Name:      teacher.Name,
					Email:     teacher.Email,
					Role:      teacher.Role,
					CreatedBy: createdBy,
				})

				if err != nil {
					session().Rollback(tx)
					return
				}
			}
		}
	}

	return session().Commit(tx)
}
