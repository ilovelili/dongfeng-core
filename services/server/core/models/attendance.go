package models

// Attendance attendance entry
type Attendance struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=attendances"`
	Date      string `dapper:"date"`
	Class     string `dapper:"class"`
	Name      string `dapper:"name"`
	CreatedBy string `dapper:"created_by"`
}
