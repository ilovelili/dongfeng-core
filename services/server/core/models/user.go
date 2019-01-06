package models

// User user profile entity
type User struct {
	ID       string      `json:"id" dapper:"id,primarykey,table=users"`
	Email    string      `json:"email" dapper:"email"`
	Name     string      `json:"name" dapper:"name"`
	Avatar   string      `json:"picture,omitempty" dapper:"avatar"`
	Setting  int64       `dapper:"settings"`
	Settings []*Settings `json:"settings" dapper:"-"`
	Role     string      `json:"role" dapper:"role"`
}

// Settings settings master
type Settings struct {
	ID      int64  `json:"id" dapper:"id,primarykey,autoincrement,table=settings"`
	Name    string `json:"name" dapper:"name"`
	Value   int64  `json:"value" dapper:"value"`
	Enabled bool   `json:"enabled"`
}
