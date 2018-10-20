package models

// Operation operation log entry
type Operation struct {
	ID        string `dapper:"uid,primarykey,table=operation_logs"`
	UserID    string `dapper:"user_id"`
	Operation string `dapper:"operation"`
	Category  string `dapper:"category"`
	Time      string `dapper:"created_at"`
}
