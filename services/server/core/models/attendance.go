package models

// Attendance attendance entry
type Attendance struct {
	ID        int    `dapper:"id,primarykey,table=attendances"`
	Date      string `dapper:"date"`
	Class     string `dapper:"class"`
	Name      string `dapper:"name"`
	CreatedBy string `dapper:"created_by"`
}
