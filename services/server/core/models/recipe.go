package models

// RecipeCombined recipe combined with nutrition
type RecipeCombined struct {
	ID                int64   `dapper:"id"`
	Name              string  `dapper:"name"`
	UnitAmount        float64 `dapper:"unit_amount"`
	Ingredient        int64   `dapper:"ingredient_id"`
	IngredientName    string  `dapper:"ingredient_name"`
	Carbohydrate      float64 `dapper:"carbohydrate"`
	Dietaryfiber      float64 `dapper:"dietaryfiber"`
	Protein           float64 `dapper:"protein"`
	Fat               float64 `dapper:"fat"`
	Heat              float64 `dapper:"heat"`
	Protein100g       float64 `dapper:"protein_100g"`
	ProteinDaily      float64 `dapper:"protein_daily"`
	Fat100g           float64 `dapper:"fat_100g"`
	FatDaily          float64 `dapper:"fat_daily"`
	Carbohydrate100g  float64 `dapper:"carbohydrate_100g"`
	CarbohydrateDaily float64 `dapper:"carbohydrate_daily"`
	Heat100g          float64 `dapper:"heat_100g"`
	HeatDaily         float64 `dapper:"heat_daily"`
	Calcium100g       float64 `dapper:"calcium_100g"`
	CalciumDaily      float64 `dapper:"calcium_daily"`
	Iron100g          float64 `dapper:"iron_100g"`
	IronDaily         float64 `dapper:"iron_daily"`
	Zinc100g          float64 `dapper:"zinc_100g"`
	ZincDaily         float64 `dapper:"zinc_daily"`
	VA100g            float64 `dapper:"va_100g"`
	VADaily           float64 `dapper:"va_daily"`
	VB1100g           float64 `dapper:"vb1_100g"`
	VB1Daily          float64 `dapper:"vb1_daily"`
	VB2100g           float64 `dapper:"vb2_100g"`
	VB2Daily          float64 `dapper:"vb2_daily"`
	VC100g            float64 `dapper:"vc_100g"`
	VCDaily           float64 `dapper:"vc_daily"`
	Category          string  `dapper:"category"`
}

// Recipe recipe
type Recipe struct {
	ID         int64   `dapper:"id,primarykey,autoincrement,table=recipes"`
	Name       string  `dapper:"name"`
	Ingredient int64   `dapper:"ingredient_id"`
	UnitAmount float64 `dapper:"unit_amount"`
	CreatedBy  string  `dapper:"created_by"`
}

// RecipeNutrition recipe nutrition
type RecipeNutrition struct {
	ID           int64   `dapper:"id,primarykey,autoincrement,table=recipe_nutritions"`
	Recipe       string  `dapper:"recipe"`
	Carbohydrate float64 `dapper:"carbohydrate"`
	Dietaryfiber float64 `dapper:"dietaryfiber"`
	Protein      float64 `dapper:"protein"`
	Fat          float64 `dapper:"fat"`
	Heat         float64 `dapper:"heat"`
}
