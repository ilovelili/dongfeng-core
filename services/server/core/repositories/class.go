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
func (r *ClassRepository) DeleteInsert(classes []*proto.Class) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	for idx, class := range classes {
		name, createdBy := class.GetName(), class.GetCreatedBy()
		if idx == 0 {
			_, err = session().ExecTx(tx, "CALL spDeleteClasses()")
			if err != nil {
				session().Rollback(tx)
				return
			}
		}

		err = session().InsertTx(tx, &models.Class{
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
