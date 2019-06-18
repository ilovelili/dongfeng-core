package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// RecipeRepository recipe repository
type RecipeRepository struct{}

// NewRecipeRepository init recipe repository
func NewRecipeRepository() *RecipeRepository {
	return &RecipeRepository{}
}

// Select get recipe by name
func (r *RecipeRepository) Select(names []string) (recipes []*models.RecipeCombined, err error) {
	var query string
	querybuilder := Table("recipes").Alias("r").
		Join("ingredients").Alias("i").On("r.ingredient_id", "i.id").Query().
		LeftOuterJoin("recipe_nutritions").Alias("rn").On("r.name", "rn.recipe").Query().
		LeftOuterJoin("ingredient_nutritions").Alias("inu").On("i.material", "inu.alias").Query().
		LeftOuterJoin("ingredient_categories").Alias("ic").On("inu.ingredient_category_id", "ic.id").
		Project(
			"r.id as id",
			"r.name as name",
			"IFNULL(r.unit_amount, 0) as unit_amount",
			"r.created_by as created_by",
			"i.id as ingredient_id",
			"i.material as ingredient_name",
			`IFNULL(rn.carbohydrate, 0) as carbohydrate`,
			`IFNULL(rn.dietaryfiber, 0) as dietaryfiber`,
			`IFNULL(rn.protein, 0) as protein`,
			`IFNULL(rn.fat, 0) as fat`,
			`IFNULL(rn.heat, 0) as heat`,
			`IFNULL(ic.category, "") as category`,
			`IFNULL(inu.protein_100g, 0) as protein_100g`,
			`IFNULL(inu.protein_daily, 0) as protein_daily`,
			`IFNULL(inu.fat_100g, 0) as fat_100g`,
			`IFNULL(inu.fat_daily, 0) as fat_daily`,
			`IFNULL(inu.carbohydrate_100g, 0) as carbohydrate_100g`,
			`IFNULL(inu.carbohydrate_daily, 0) as carbohydrate_daily`,
			`IFNULL(inu.heat_100g, 0) as heat_100g`,
			`IFNULL(inu.heat_daily, 0) as heat_daily`,
			`IFNULL(inu.calcium_100g, 0) as calcium_100g`,
			`IFNULL(inu.calcium_daily, 0) as calcium_daily`,
			`IFNULL(inu.iron_100g, 0) as iron_100g`,
			`IFNULL(inu.iron_daily, 0) as iron_daily`,
			`IFNULL(inu.zinc_100g, 0) as zinc_100g`,
			`IFNULL(inu.zinc_daily, 0) as zinc_daily`,
			`IFNULL(inu.va_100g, 0) as va_100g`,
			`IFNULL(inu.va_daily, 0) as va_daily`,
			`IFNULL(inu.vb1_100g, 0) as vb1_100g`,
			`IFNULL(inu.vb1_daily, 0) as vb1_daily`,
			`IFNULL(inu.vb2_100g, 0) as vb2_100g`,
			`IFNULL(inu.vb2_daily, 0) as vb2_daily`,
			`IFNULL(inu.vc_100g, 0) as vc_100g`,
			`IFNULL(inu.vc_daily, 0) as vc_daily`,
		)

	if len(names) != 0 {
		if len(names) == 1 && names[0] == "" {
			query = querybuilder.Sql()
		} else {
			query = querybuilder.Where().In("r.name", names).Sql()
		}
	} else {
		query = querybuilder.Sql()
	}

	// no rows is actually not an error
	if err = session().Find(query, nil).All(&recipes); err != nil && norows(err) {
		err = nil
	}

	return
}

// Update update recipe
func (r *RecipeRepository) Update(recipe *models.Recipe) (err error) {
	query := Table("recipes").Alias("r").
		Project(
			"r.id as id",
			"r.name as name",
			"r.ingredient_id as ingredient_id",
			"IFNULL(r.unit_amount, 0) as unit_amount",
			"r.created_by as created_by",
		).
		Where().Eq("r.id", recipe.ID).Sql()

	var _recipe models.Recipe
	err = session().Find(query, nil).Single(&_recipe)
	if err != nil || _recipe.ID == 0 {
		return fmt.Errorf("recipe row doesnot exist")
	}

	_recipe.UnitAmount = recipe.UnitAmount
	return session().Update(_recipe)
}
