package models

// Holiday holiday entity
type Holiday struct {
	ID   int64  `dapper:"id,primarykey,autoincrement,table=holidays"`
	From string `dapper:"from"`
	To   string `dapper:"to"`
}

// HolidayType 0: working day | 1: weekend | 2: holiday
type HolidayType struct {
	Date string
	Type int64
}
