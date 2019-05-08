package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// MenuController menu controller
type MenuController struct {
	repository *repositories.MenuRepository
}

// NewMenuController new controller
func NewMenuController() *MenuController {
	return &MenuController{
		repository: repositories.NewMenuRepository(),
	}
}

// GetMenus get menus
func (c *MenuController) GetMenus(from string, to string, mealid int64, seniororjunior int64) (menus []*models.Menu, err error) {
	return c.repository.Select(from, to, mealid, seniororjunior)
}
