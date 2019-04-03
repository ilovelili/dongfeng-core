package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ClassRepository class repository
type ClassRepository struct{}

// NewClassRepository init class repository
func NewClassRepository() *ClassRepository {
	return &ClassRepository{}
}

// Select select Class
func (r *ClassRepository) Select() (classes []*models.Class, err error) {
	query := Table("classes").Sql()
	err = session().Find(query, nil).All(&classes)
	return
}

// DeleteInsert delete insert Class
func (r *ClassRepository) DeleteInsert(classes []*proto.ClassItem) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for idx, class := range classes {
		createdBy := class.GetCreatedBy()

		if idx == 0 {
			_, err = session().ExecTx(tx, "CALL spDeleteClasses()")
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		err = session().InsertTx(tx, &models.Class{
			Name:      class.Name,
			CreatedBy: createdBy,
		})

		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
