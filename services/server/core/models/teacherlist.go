package models

// Teacherlist teacher list entity
type Teacherlist struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=teacherlists"`
	Year      string `dapper:"year"`
	Name      string `dapper:"name"`
	Class     string `dapper:"class"`
	Email     string `dapper:"email"`
	Role      string `dapper:"role"`
	CreatedBy string `dapper:"created_by"`
}
