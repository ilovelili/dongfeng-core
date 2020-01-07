package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// RoleController Role controller
type RoleController struct {
	repository *repositories.RoleRepository
}

// NewRoleController new controller
func NewRoleController() *RoleController {
	return &RoleController{
		repository: repositories.NewRoleRepository(),
	}
}

// GetRole get role
func (c *RoleController) GetRole(user string) (string, error) {
	role, err := c.repository.Select(user)
	if err != nil {
		return "", err
	}
	return role.Role, nil
}

// SaveRole save role
func (c *RoleController) SaveRole(role *models.Role) error {
	return c.repository.Upsert(role)
}
