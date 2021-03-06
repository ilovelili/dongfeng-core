package handlers

import (
	"context"
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

type recipemapvalue struct {
	Ingredients  []string
	Carbohydrate float64
	Dietaryfiber float64
	Protein      float64
	Fat          float64
	Heat         float64
}

// GetRecipes get recipes
func (f *Facade) GetRecipes(ctx context.Context, req *proto.GetRecipeRequest, rsp *proto.GetRecipeResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	recipecontroller := controllers.NewRecipeController()
	recipes, err := recipecontroller.SelectRecipes(req.GetNames())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetRecipe)
	}

	recipemap := make(map[string] /*recipe name*/ *proto.Recipe)
	for _, recipe := range recipes {
		ingredient := &proto.Ingredient{
			Ingredient:        recipe.IngredientName,
			Protein_100G:      recipe.Protein100g,
			ProteinDaily:      recipe.ProteinDaily,
			Fat_100G:          recipe.Fat100g,
			FatDaily:          recipe.FatDaily,
			Carbohydrate_100G: recipe.Carbohydrate100g,
			CarbohydrateDaily: recipe.CarbohydrateDaily,
			Heat_100G:         recipe.Heat100g,
			HeatDaily:         recipe.HeatDaily,
			Calcium_100G:      recipe.Calcium100g,
			CalciumDaily:      recipe.CalciumDaily,
			Iron_100G:         recipe.Iron100g,
			IronDaily:         recipe.IronDaily,
			Zinc_100G:         recipe.Zinc100g,
			ZincDaily:         recipe.ZincDaily,
			Va_100G:           recipe.VA100g,
			VaDaily:           recipe.VADaily,
			Vb1_100G:          recipe.VB1100g,
			Vb1Daily:          recipe.VB1Daily,
			Vb2_100G:          recipe.VB2100g,
			Vb2Daily:          recipe.VB2Daily,
			Vc_100G:           recipe.VC100g,
			VcDaily:           recipe.VCDaily,
			Category:          recipe.Category,
		}

		if v, ok := recipemap[recipe.Name]; ok {
			recipemap[recipe.Name].Ingredients = append(v.GetIngredients(), ingredient)
		} else {
			recipemap[recipe.Name] = &proto.Recipe{
				Recipe:       recipe.Name,
				Ingredients:  []*proto.Ingredient{ingredient},
				Carbohydrate: recipe.Carbohydrate,
				Dietaryfiber: recipe.Dietaryfiber,
				Protein:      recipe.Protein,
				Fat:          recipe.Fat,
				Heat:         recipe.Heat,
			}
		}
	}

	r := []*proto.Recipe{}
	for _, v := range recipemap {
		r = append(r, &proto.Recipe{
			Recipe:       v.Recipe,
			Ingredients:  v.Ingredients,
			Carbohydrate: v.Carbohydrate,
			Dietaryfiber: v.Dietaryfiber,
			Protein:      v.Protein,
			Fat:          v.Fat,
			Heat:         v.Heat,
		})
	}

	rsp.Recipes = r
	return nil
}

// UpdateRecipe update recipe
func (f *Facade) UpdateRecipe(ctx context.Context, req *proto.UpdateRecipeRequest, rsp *proto.UpdateRecipeResponse) error {
	return fmt.Errorf("Not implemented")
}
