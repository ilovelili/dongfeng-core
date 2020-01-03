package models

// Role role entity
type Role struct {
	Role string `dapper:"role"`
	Path string `dapper:"path"`
}
