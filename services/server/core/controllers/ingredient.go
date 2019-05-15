package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
)

// IngredientController ingredient controller
type IngredientController struct {
	repository *repositories.IngredientRepository
}

// NewIngredientController new controller
func NewIngredientController() *IngredientController {
	return &IngredientController{
		repository: repositories.NewIngredientRepository(),
	}
}

// SelectIngredientNutritions select ingredient nutritions
func (c *IngredientController) SelectIngredientNutritions(names []string) (ingredients []*models.IngredientNutrition, err error) {
	ingredients, err = c.repository.SelectIngredientNutritions(names)
	if err != nil {
		return
	}

	// get categories master
	categories, caterr := c.repository.SelectIngredientCategories()
	if caterr != nil {
		err = caterr
		return
	}

	for _, ingredient := range ingredients {
		for _, category := range categories {
			if ingredient.CategoryID == category.ID {
				ingredient.Category = category.Category
				break
			}
		}
	}

	return
}

// SaveIngredientNutritions save ingredient nutritions
func (c *IngredientController) SaveIngredientNutritions(ingredients []*models.IngredientNutrition) error {
	for _, ingredient := range ingredients {
		id, err := c.repository.SelectIngredientCategory(ingredient.Category)
		if err != nil {
			return utils.NewError(errorcode.CoreInvalidIngredientCategory)
		}

		ingredient.CategoryID = id
	}

	err := c.repository.UpsertIngredientNutritions(ingredients)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveIngredient)
	}

	return nil
}
