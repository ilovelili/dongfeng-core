package models

// Attendance attendance entry
type Attendance struct {
	ID         int64  `dapper:"id,primarykey,autoincrement,table=attendances"`
	Year       string `dapper:"year"`
	Date       string `dapper:"date"`
	Class      string `dapper:"class"`
	Name       string `dapper:"name"`
	CreatedBy  string `dapper:"created_by"`
	Attendance string `dapper:"-"` // attendance flag
}
