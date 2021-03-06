package models

// Absence absence entity
type Absence struct {
	ID        int64  `dapper:"id,primarykey,autoincrement,table=absences"`
	Year      string `dapper:"year"`
	Date      string `dapper:"date"`
	Class     string `dapper:"class"`
	Name      string `dapper:"name"`
	CreatedBy string `dapper:"created_by"`
}

// Attendance attendence entity
type Attendance struct {
	Year           string
	Date           string
	Class          string
	Name           string
	AttendanceFlag bool
}

// Attendances attendances with holiday
type Attendances struct {
	Attendances []*Attendance
	Holidays    []*HolidayType
}

// AttendanceCount attendance count entity
type AttendanceCount struct {
	Sum                   int64
	JuniorSum             int64
	MiddleSum             int64
	SeniorSum             int64
	ClassAttendanceCounts []*ClassAttendanceCount
}

// ClassAttendanceCount attendance count per class
type ClassAttendanceCount struct {
	Date  string
	Class string
	Count int64
}

/*
// Attendances alias of attendance slice for sort implementaion
type Attendances []*Attendance

func (a Attendances) Len() int {
	attendances := []*Attendance(a)
	return len(attendances)
}

func (a Attendances) Swap(i, j int) {
	attendances := []*Attendance(a)
	attendances[i], attendances[j] = attendances[j], attendances[i]
}

func (a Attendances) Less(i, j int) bool {
	attendances := []*Attendance(a)
	return attendances[i].Date < attendances[j].Date
}
*/
