package models

// Namelist namelist entity
type Namelist struct {
	ID        int64  `dapper:"id,primarykey,table=namelists"`
	Year      string `dapper:"year"`
	Name      string `dapper:"name"`
	Class     string `dapper:"class_id"`
	CreatedBy string `dapper:"created_by"`
}
