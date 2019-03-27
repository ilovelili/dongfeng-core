package models

// Namelist namelist entity
type Namelist struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=namelists"`
	Year      string `dapper:"year"`
	Name      string `dapper:"name"`
	Class     string `dapper:"class"`
	CreatedBy string `dapper:"created_by"`
}
