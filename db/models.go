package db

type Student struct {
	ID    int16
	FName string
	LName string
	DOB   string
}

type Course struct {
	ID          int16
	CourseName  string
	TotalCredit int32
}

type Enroll struct {
	CourseID  int16
	StudentID int16
	Semester  string
	Score     float32
}
