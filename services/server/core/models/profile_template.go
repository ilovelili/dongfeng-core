package models

// ProfileTemplate  profile template
type ProfileTemplate struct {
	Name      string `dapper:"name,primarykey,autoincrement,table=profile_templates"`
	Profile   string `dapper:"profile"`
	Enabled   bool   `dapper:"enabled"`
	CreatedBy string `dapper:"created_by"`
}
