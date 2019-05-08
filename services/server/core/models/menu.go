package models

// Menu menu entry
type Menu struct {
	ID               int64  `dapper:"id,primarykey,autoincrement,table=menus"`
	Date             string `dapper:"date"`
	Recipe           string `dapper:"recipe"`
	BreakfastOrLunch int64  `dapper:"breakfast_or_lunch"`
	JuniorOrSenior   int64  `dapper:"junior_or_senior_class"`
}
