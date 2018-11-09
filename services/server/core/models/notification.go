package models

// Notification notification entry
type Notification struct {
	ID         int32  `dapper:"uid,primarykey,table=notifications"`
	UserID     string `dapper:"user_id"`
	CustomCode string `dapper:"custom_code"`
	Details    string `dapper:"details"`
	Link       string `dapper:"link"`
	Category   string `dapper:"category"`
	Time       string `dapper:"created_at"`
}
