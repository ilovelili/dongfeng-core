package models

// Teacher teacher entity
type Teacher struct {
	ID        int64   `dapper:"id,primarykey,autoincrement,table=teachers"`
	Year      string  `dapper:"year"`
	Name      string  `dapper:"name"`
	Class     string  `dapper:"class"`
	Email     string  `dapper:"email"`
	Role      *string `dapper:"role"`
	CreatedBy string  `dapper:"created_by"`
}

// TeacherWithoutRole teacher without role
type TeacherWithoutRole struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=teachers"`
	Year      string `dapper:"year"`
	Name      string `dapper:"name"`
	Class     string `dapper:"class"`
	Email     string `dapper:"email"`
	CreatedBy string `dapper:"created_by"`
}

// RemoveRole remove role since there is no role column in teachers table
func (t *Teacher) RemoveRole() *TeacherWithoutRole {
	return &TeacherWithoutRole{
		ID:        t.ID,
		Year:      t.Year,
		Name:      t.Name,
		Class:     t.Class,
		Email:     t.Email,
		CreatedBy: t.CreatedBy,
	}
}
