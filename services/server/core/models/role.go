package models

// Role role entity
type Role struct {
	User string `dapper:"user,primarykey,table=roles"`
	Role string `dapper:"role"`
}
