package models

// Pupil pupil entity
type Pupil struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=pupils"`
	Year      string `dapper:"year"`
	Name      string `dapper:"name"`
	Class     string `dapper:"class"`
	CreatedBy string `dapper:"created_by"`
}
