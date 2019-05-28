package models

// Profile profile entity
type Profile struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=profiles"`
	Year      string `dapper:"year"`
	Class     string `dapper:"class"`
	Name      string `dapper:"name"`
	Date      string `dapper:"date"`
	Profile   string `dapper:"profile"`
	CreatedBy string `dapper:"created_by"`
}
