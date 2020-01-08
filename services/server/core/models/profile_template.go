package models

// ProfileTemplate  profile template
type ProfileTemplate struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=profile_templates"`
	Name      string `dapper:"name"`
	Profile   string `dapper:"profile"`
	Enabled   bool   `dapper:"enabled"`
	CreatedBy string `dapper:"created_by"`
}
