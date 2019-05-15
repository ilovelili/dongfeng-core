package models

// IngredientNutrition ingredient nutrition with category
type IngredientNutrition struct {
	ID                int64   `dapper:"id,primarykey,autoincrement,table=ingredient_nutritions"`
	Ingredient        string  `dapper:"ingredient"`
	Alias             string  `dapper:"alias"`
	Category          string  `dapper:"-"`
	CategoryID        int64   `dapper:"ingredient_category_id"`
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
}

// IngredientCategory ingredient category
type IngredientCategory struct {
	ID       int64  `dapper:"id,primarykey,autoincrement,table=ingredient_categories"`
	Category string `dapper:"category"`
}
