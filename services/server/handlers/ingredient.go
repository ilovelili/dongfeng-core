package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetIngredients get ingredients
func (f *Facade) GetIngredients(ctx context.Context, req *proto.GetIngredientRequest, rsp *proto.GetIngredientResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// Unmarshal user info
	userinfo, _ := json.Marshal(claims)
	var user *models.User
	err = json.Unmarshal(userinfo, &user)

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	ingredientcontroller := controllers.NewIngredientController()
	ingredients, err := ingredientcontroller.SelectIngredientNutritions(req.GetIngredients())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetIngredient)
	}

	_ingredients := []*proto.Ingredient{}
	for _, ingredient := range ingredients {
		_ingredients = append(_ingredients, &proto.Ingredient{
			Ingredient:        ingredient.Ingredient,
			Alias:             ingredient.Alias,
			Protein_100G:      ingredient.Protein100g,
			ProteinDaily:      ingredient.ProteinDaily,
			Fat_100G:          ingredient.Fat100g,
			FatDaily:          ingredient.FatDaily,
			Carbohydrate_100G: ingredient.Carbohydrate100g,
			CarbohydrateDaily: ingredient.CarbohydrateDaily,
			Heat_100G:         ingredient.Heat100g,
			HeatDaily:         ingredient.HeatDaily,
			Calcium_100G:      ingredient.Calcium100g,
			CalciumDaily:      ingredient.CalciumDaily,
			Iron_100G:         ingredient.Iron100g,
			IronDaily:         ingredient.IronDaily,
			Zinc_100G:         ingredient.Zinc100g,
			ZincDaily:         ingredient.ZincDaily,
			Va_100G:           ingredient.VA100g,
			VaDaily:           ingredient.VADaily,
			Vb1_100G:          ingredient.VB1100g,
			Vb1Daily:          ingredient.VB1Daily,
			Vb2_100G:          ingredient.VB2100g,
			Vb2Daily:          ingredient.VB2Daily,
			Vc_100G:           ingredient.VC100g,
			VcDaily:           ingredient.VCDaily,
			Category:          ingredient.Category,
		})
	}

	rsp.Ingredients = _ingredients
	return nil
}

// UpdateIngredient update ingredient
func (f *Facade) UpdateIngredient(ctx context.Context, req *proto.UpdateIngredientRequest, rsp *proto.UpdateIngredientResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// Unmarshal user info
	userinfo, _ := json.Marshal(claims)
	var user *models.User
	err = json.Unmarshal(userinfo, &user)

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	ingredientcontroller := controllers.NewIngredientController()
	ingredients := []*models.IngredientNutrition{}
	for _, ingredient := range req.GetIngredients() {
		ingredients = append(ingredients, &models.IngredientNutrition{
			Ingredient:        ingredient.GetIngredient(),
			Alias:             ingredient.GetAlias(),
			Category:          ingredient.GetCategory(),
			Protein100g:       ingredient.GetProtein_100G(),
			ProteinDaily:      ingredient.GetProteinDaily(),
			Fat100g:           ingredient.GetFat_100G(),
			FatDaily:          ingredient.GetFatDaily(),
			Carbohydrate100g:  ingredient.GetCarbohydrate_100G(),
			CarbohydrateDaily: ingredient.GetCarbohydrateDaily(),
			Heat100g:          ingredient.GetHeat_100G(),
			HeatDaily:         ingredient.GetHeatDaily(),
			Calcium100g:       ingredient.GetCalcium_100G(),
			CalciumDaily:      ingredient.GetCalciumDaily(),
			Iron100g:          ingredient.GetIron_100G(),
			IronDaily:         ingredient.GetIronDaily(),
			Zinc100g:          ingredient.GetZinc_100G(),
			ZincDaily:         ingredient.GetZincDaily(),
			VA100g:            ingredient.GetVa_100G(),
			VADaily:           ingredient.GetVaDaily(),
			VB1100g:           ingredient.GetVb1_100G(),
			VB1Daily:          ingredient.GetVb1Daily(),
			VB2100g:           ingredient.GetVb2_100G(),
			VB2Daily:          ingredient.GetVb2Daily(),
			VC100g:            ingredient.GetVc_100G(),
			VCDaily:           ingredient.GetVcDaily(),
		})
	}

	if err = ingredientcontroller.SaveIngredientNutritions(ingredients); err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveIngredient)
	}

	f.syslog(notification.IngredientUpdated(user.ID))
	return nil
}

// UpdateIngredients update ingredients
func (f *Facade) UpdateIngredients(ctx context.Context, req *proto.UpdateIngredientRequest, rsp *proto.UpdateIngredientResponse) error {
	return fmt.Errorf("not implemented")
}
