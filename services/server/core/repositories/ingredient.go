package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// IngredientRepository ingredient repository
type IngredientRepository struct{}

// NewIngredientRepository init ingredient repository
func NewIngredientRepository() *IngredientRepository {
	return &IngredientRepository{}
}

// SelectIngredientCategory select ingredient category id
func (r *IngredientRepository) SelectIngredientCategory(category string) (id int64, err error) {
	query := Table("ingredient_categories").Alias("i").Project("i.id").Where().Eq("i.category", category).Sql()
	err = session().Find(query, nil).Scalar(&id)
	return
}

// SelectIngredientCategories select all ingredient categories
func (r *IngredientRepository) SelectIngredientCategories() (categories []*models.IngredientCategory, err error) {
	query := Table("ingredient_categories").Sql()
	err = session().Find(query, nil).All(&categories)
	return
}

// SelectIngredientNutritions select ingredient nutritions
func (r *IngredientRepository) SelectIngredientNutritions(names []string) (ingredients []*models.IngredientNutrition, err error) {
	querybuilder := Table("ingredient_nutritions").Alias("inu")

	var query string
	if len(names) != 0 {
		if len(names) == 1 && names[0] == "" {
			query = querybuilder.Sql()
		} else {
			query = querybuilder.Where().In("inu.alias", names).Sql()
		}
	} else {
		query = querybuilder.Sql()
	}

	if err = session().Find(query, nil).All(&ingredients); err != nil && norows(err) {
		err = nil
	}

	return
}

// UpsertIngredientNutritions upsert ingredient nutrition
func (r *IngredientRepository) UpsertIngredientNutritions(ingredients []*models.IngredientNutrition) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// upsert by loop
	for _, ingredient := range ingredients {
		query := Table("ingredient_nutritions").Alias("i").Where().Eq("i.ingredient", ingredient.Ingredient).Sql()

		var i models.IngredientNutrition
		err := session().Find(query, nil).Single(&i)
		if err != nil || 0 == i.ID {
			err = session().InsertTx(tx, ingredient)
		} else {
			ingredient.ID = i.ID
			err = session().UpdateTx(tx, ingredient)
		}

		if err != nil {
			session().Rollback(tx)
			return err
		}
	}

	return session().Commit(tx)
}
