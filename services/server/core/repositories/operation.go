package repositories

import (
	"github.com/ilovelili/dongfeng/core/services/server/core/models"
)

// OperationRepository friends repository
type OperationRepository struct{}

// NewOperationRepository init UserProfile repository
func NewOperationRepository() *OperationRepository {
	return &OperationRepository{}
}

// SelectOperations select operation logs
func (r *OperationRepository) SelectOperations(uid string, adminonly bool) (operations []*models.Operation, err error) {
	clause := Table("operation_logs").Alias("o").
		Join("categories").Alias("c").On("o.category_id", "c.id").
		Project("o.id", "o.user_id", "o.operation", "o.created_at", "c.description").
		Where().Eq("o.user_id", uid)

	if !adminonly {
		clause = clause.Eq("c.admin_only", 0)
	}

	query := clause.Order().Desc("o.created_at").Sql()
	// no rows is actually not an error
	if err = session().Find(query, nil).All(&operations); err != nil && norows(err) {
		err = nil
	}

	return
}

// Insert insert operation
func (r *OperationRepository) Insert(operation *models.Operation) error {
	return insertTx(operation)
}
