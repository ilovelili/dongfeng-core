package controllers

import (
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

// GetAccessiblePaths get accessible paths
func (c *RoleController) GetAccessiblePaths(user string) ([]string, error) {
	return c.repository.Select(user)
}
