package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// OperationController user profile controller
type OperationController struct {
	repository *repositories.OperationRepository
}

// NewOperationController new controller
func NewOperationController() *OperationController {
	return &OperationController{
		repository: repositories.NewOperationRepository(),
	}
}

// GetOperations get operations
func (c *OperationController) GetOperations(uid string, adminonly bool) ([]*models.Operation, error) {
	return c.repository.SelectOperations(uid, adminonly)
}

// Save save operations
func (c *OperationController) Save(operation *models.Operation) error {
	return c.repository.Insert(operation)
}
