package models

// User user profile entity
type User struct {
	ID       string      `json:"ID" dapper:"id,primarykey,table=users"`
	Email    string      `json:"Email" dapper:"email"`
	Name     string      `json:"Username" dapper:"name"`
	Avatar   string      `json:"Photo,omitempty" dapper:"avatar"`
	Setting  int64       `dapper:"settings"`
	Settings []*Settings `json:"settings" dapper:"-"`	
}

// Settings settings master
type Settings struct {
	ID      int64  `json:"id" dapper:"id,primarykey,autoincrement,table=settings"`
	Name    string `json:"name" dapper:"name"`
	Value   int64  `json:"value" dapper:"value"`
	Enabled bool   `json:"enabled"`
}