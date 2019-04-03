package models

// Class class list entity
type Class struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=classes"`
	Name      string `dapper:"name"`
	CreatedBy string `dapper:"created_by"`
}
