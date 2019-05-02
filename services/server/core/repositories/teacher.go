package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherRepository teacher repository
type TeacherRepository struct{}

// NewTeacherRepository init teacher repository
func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{}
}

// Select select teachers
func (r *TeacherRepository) Select(class, year string) (teachers []*models.Teacher, err error) {
	querybuilder := Table("teachers").Alias("t").Where()
	var query string

	if class == "" && year == "" {
		querybuilder = querybuilder.Eq("1", "1")
	} else {
		if class != "" {
			querybuilder = querybuilder.Like("t.class", fmt.Sprintf("%%%s%%", class))
		}

		if year != "" {
			querybuilder = querybuilder.Eq("t.year", year)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&teachers)
	return
}

// Update update teacher
func (r *TeacherRepository) Update(teacher *models.Teacher) (err error) {
	return session().Update(teacher)
}

// DeleteInsert delete insert teachers
func (r *TeacherRepository) DeleteInsert(teachers []*proto.Teacher) (err error) {
	teachersmap := make(map[string][]*proto.Teacher)
	for _, teacher := range teachers {
		key := teacher.Year
		if v, ok := teachersmap[key]; !ok {
			teachersmap[key] = []*proto.Teacher{teacher}
		} else {
			teachersmap[key] = append(v, teacher)
		}
	}

	tx, err := session().Begin()
	if err != nil {
		return
	}

	for year := range teachersmap {
		_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeleteTeachers('%s')", year))
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	for _, teacher := range teachers {
		err = session().InsertTx(tx, &models.Teacher{
			Year:      teacher.GetYear(),
			Class:     teacher.GetClass(),
			Name:      teacher.GetName(),
			Email:     teacher.GetEmail(),
			Role:      teacher.GetRole(),
			CreatedBy: teacher.GetCreatedBy(),
		})

		if err != nil {
			session().Rollback(tx)
			return
		}

	}

	return session().Commit(tx)
}
