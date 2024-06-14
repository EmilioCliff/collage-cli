package db

import (
	"database/sql"
)

func AddStudent(arg Student) (Student, error) {
	query := `
		INSERT INTO students 
		(f_name, l_name, dob)  
		VALUES(?, ?, ?)
		RETURNING id, f_name, l_name, dob
	`

	row := DB.QueryRow(query, arg.FName, arg.LName, arg.DOB)
	var result Student
	if err := row.Scan(&result.ID, &result.FName, &result.LName, &result.DOB); err != nil {
		return Student{}, err
	}

	return result, nil
}

func ListStudents() ([]Student, error) {
	query := `
		SELECT * FROM students ORDER BY f_name
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Student
	for rows.Next() {
		var student Student
		if err = rows.Scan(&student.ID, &student.FName, &student.LName, &student.DOB); err != nil {
			return nil, err
		}

		result = append(result, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetStudent(studentID int16) (Student, error) {
	query := `
		SELECT * FROM students
		WHERE id = ?
	`
	row := DB.QueryRow(query, studentID)

	var result Student
	if err := row.Scan(&result.ID, &result.FName, &result.LName, &result.DOB); err != nil {
		return Student{}, err
	}

	return result, nil
}

type CourseResponse struct {
	CourseName string
	Semester   string
}

type searchStudentResponse struct {
	ID            int16
	FName         string
	LName         string
	DOB           string
	Courses       []CourseResponse
	OverrallScore float32
	LetterGrade   string
}

func SearchStudentByID(id int16) ([]searchStudentResponse, error) {
	query := `
		SELECT * FROM students
		WHERE id = ?
		ORDER BY f_name
	`
	rows, err := DB.Query(query, id)
	if err != nil {
		return []searchStudentResponse{}, err
	}
	defer rows.Close()

	result, err := rowsCompute(rows)
	if err != nil {
		return []searchStudentResponse{}, err
	}

	return result, nil
}

// SELECT * FROM students WHERE LOWER(firstname) LIKE LOWER(?) || '%';
func SearchStudentByLName(lName string) ([]searchStudentResponse, error) {
	query := `
		SELECT * FROM students
		WHERE LOWER(l_name) LIKE LOWER(?) || '%'
		ORDER BY f_name;
	`
	rows, err := DB.Query(query, lName)
	if err != nil {
		return []searchStudentResponse{}, err
	}
	defer rows.Close()

	result, err := rowsCompute(rows)
	if err != nil {
		return []searchStudentResponse{}, err
	}

	return result, nil
}

func rowsCompute(rows *sql.Rows) ([]searchStudentResponse, error) {
	var result []searchStudentResponse
	var err error
	for rows.Next() {
		var response searchStudentResponse
		var student Student
		if err = rows.Scan(&student.ID, &student.FName, &student.LName, &student.DOB); err != nil {
			return []searchStudentResponse{}, err
		}

		response.OverrallScore, err = GetStudentAverageScore(student.ID)
		if err != nil {
			return []searchStudentResponse{}, err
		}

		response.Courses, err = GetStudentsCourse(student.ID)
		if err != nil {
			return []searchStudentResponse{}, err
		}

		response.ID = student.ID
		response.FName = student.FName
		response.LName = student.LName
		response.DOB = student.DOB
		response.LetterGrade = SelectGrade(response.OverrallScore)

		result = append(result, response)
	}

	if err = rows.Err(); err != nil {
		return []searchStudentResponse{}, err
	}

	if err = rows.Close(); err != nil {
		return []searchStudentResponse{}, err
	}

	return result, nil
}

func SelectGrade(score float32) string {
	var grade string

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	return grade
}
