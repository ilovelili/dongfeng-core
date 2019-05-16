package repositories

import (
	"fmt"

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

// SelectIngredientNames select ingredient names
func (r *IngredientRepository) SelectIngredientNames(pattern string) (ingredients []*models.IngredientNutrition, err error) {
	query := Table("ingredient_nutritions").Alias("inu").Where().Like("inu.ingredient", fmt.Sprintf("%%%s%%", pattern)).Sql()
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

			if i.Alias != "" && ingredient.Alias == "" {
				ingredient.Alias = i.Alias
			}

			if i.Calcium100g != 0 && ingredient.Calcium100g == 0 {
				ingredient.Calcium100g = i.Calcium100g
			}
			if i.CalciumDaily != 0 && ingredient.CalciumDaily == 0 {
				ingredient.CalciumDaily = i.CalciumDaily
			}

			if i.Carbohydrate100g != 0 && ingredient.Carbohydrate100g == 0 {
				ingredient.Carbohydrate100g = i.Carbohydrate100g
			}
			if i.CarbohydrateDaily != 0 && ingredient.CarbohydrateDaily == 0 {
				ingredient.CarbohydrateDaily = i.CarbohydrateDaily
			}

			if i.CategoryID != 0 && ingredient.CategoryID == 0 {
				ingredient.CategoryID = i.CategoryID
			}

			if i.Fat100g != 0 && ingredient.Fat100g == 0 {
				ingredient.Fat100g = i.Fat100g
			}
			if i.FatDaily != 0 && ingredient.FatDaily == 0 {
				ingredient.FatDaily = i.FatDaily
			}

			if i.Heat100g != 0 && ingredient.Heat100g == 0 {
				ingredient.Heat100g = i.Heat100g
			}
			if i.HeatDaily != 0 && ingredient.HeatDaily == 0 {
				ingredient.HeatDaily = i.HeatDaily
			}

			if i.Ingredient != "" && ingredient.Ingredient == "" {
				ingredient.Ingredient = i.Ingredient
			}

			if i.Iron100g != 0 && ingredient.Iron100g == 0 {
				ingredient.Iron100g = i.Iron100g
			}
			if i.IronDaily != 0 && ingredient.IronDaily == 0 {
				ingredient.IronDaily = i.IronDaily
			}

			if i.Protein100g != 0 && ingredient.Protein100g == 0 {
				ingredient.Protein100g = i.Protein100g
			}
			if i.ProteinDaily != 0 && ingredient.ProteinDaily == 0 {
				ingredient.ProteinDaily = i.ProteinDaily
			}

			if i.VA100g != 0 && ingredient.VA100g == 0 {
				ingredient.VA100g = i.VA100g
			}
			if i.VADaily != 0 && ingredient.VADaily == 0 {
				ingredient.VADaily = i.VADaily
			}

			if i.VB1100g != 0 && ingredient.VB1100g == 0 {
				ingredient.VB1100g = i.VB1100g
			}
			if i.VB1Daily != 0 && ingredient.VB1Daily == 0 {
				ingredient.VB1Daily = i.VB1Daily
			}

			if i.VB2100g != 0 && ingredient.VB2100g == 0 {
				ingredient.VB2100g = i.VB2100g
			}
			if i.VB2Daily != 0 && ingredient.VB2Daily == 0 {
				ingredient.VB2Daily = i.VB2Daily
			}

			if i.VC100g != 0 && ingredient.VC100g == 0 {
				ingredient.VC100g = i.VC100g
			}
			if i.VCDaily != 0 && ingredient.VCDaily == 0 {
				ingredient.VCDaily = i.VCDaily
			}

			if i.Zinc100g != 0 && ingredient.Zinc100g == 0 {
				ingredient.Zinc100g = i.Zinc100g
			}
			if i.ZincDaily != 0 && ingredient.ZincDaily == 0 {
				ingredient.ZincDaily = i.ZincDaily
			}

			err = session().UpdateTx(tx, ingredient)
		}

		if err != nil {
			session().Rollback(tx)
			return err
		}
	}

	return session().Commit(tx)
}
