package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// RecipeController recipe controller
type RecipeController struct {
	repository *repositories.RecipeRepository
}

// NewRecipeController new controller
func NewRecipeController() *RecipeController {
	return &RecipeController{
		repository: repositories.NewRecipeRepository(),
	}
}

// SelectRecipes select recipes
func (c *RecipeController) SelectRecipes(names []string) ([]*models.RecipeCombined, error) {
	return c.repository.Select(names)
}
