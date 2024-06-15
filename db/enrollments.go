package db

import (
	"database/sql"
	"fmt"
)

func EnrollStudent(arg Enroll) (Enroll, error) {
	query := `
	INSERT INTO enrollment
	(course_id, student_id, semester, score)
	VALUES (?, ?, ?, ?)
	RETURNING course_id, student_id, semester, score
`
	row := DB.QueryRow(query, arg.CourseID, arg.StudentID, arg.Semester, arg.Score)

	var result Enroll
	if err := row.Scan(&result.CourseID, &result.StudentID, &result.Semester, &arg.Score); err != nil {
		return Enroll{}, err
	}

	return result, nil
}

func ListEnrollments() ([]Enroll, error) {
	query := `
		SELECT * FROM enrollment
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Enroll
	for rows.Next() {
		var enroll Enroll
		if err = rows.Scan(&enroll.CourseID, &enroll.StudentID, &enroll.Semester, &enroll.Score); err != nil {
			return nil, err
		}

		result = append(result, enroll)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetStudentAverageScore(studentID int16) (float32, error) {
	query := `
		SELECT  AVG(score) AS average_score 
		FROM enrollment
		WHERE student_id = ?
	`
	row := DB.QueryRow(query, studentID)
	var score sql.NullFloat64
	if err := row.Scan(&score); err != nil {
		return 0, err
	}

	if score.Valid {
		return float32(score.Float64), nil
	}

	return 0, nil
}

func GetStudentsCourse(studentID int16) ([]CourseResponse, error) {
	query := `
		SELECT course_id, semester FROM enrollment
		WHERE student_id = ?
	`
	rows, err := DB.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []CourseResponse
	for rows.Next() {
		var response CourseResponse
		var courseID int16
		err := rows.Scan(&courseID, &response.Semester)
		if err != nil {
			return nil, err
		}

		course, err := GetCourse(courseID)
		if err != nil {
			return nil, err
		}

		response.CourseName = course.CourseName

		courses = append(courses, response)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return courses, nil
}

func GetCourseAverageScore(courseID int16) (float32, error) {
	query := `
		SELECT  AVG(score) AS average_score 
		FROM enrollment
		WHERE course_id = ?
	`
	row := DB.QueryRow(query, courseID)
	var score sql.NullFloat64
	if err := row.Scan(&score); err != nil {
		return 0, err
	}

	if score.Valid {
		return float32(score.Float64), nil
	}

	return 0, nil
}

func GetCourseAndSemesterAverageScore(courseID int16, semester string) (float32, error) {
	query := `
		SELECT  AVG(score) AS average_score 
		FROM enrollment
		WHERE course_id = ? AND semester = ?
	`
	row := DB.QueryRow(query, courseID, semester)
	var score sql.NullFloat64
	if err := row.Scan(&score); err != nil {
		return 0, err
	}

	if score.Valid {
		return float32(score.Float64), nil
	}

	return 0, nil
}

func GetCourseEnrollments(courseId int16) ([]StudentResponse, error) {
	query := `
		SELECT student_id, semester FROM enrollment
		WHERE course_id = ?
	`
	rows, err := DB.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []StudentResponse
	for rows.Next() {
		var response StudentResponse
		var studentID int16
		err := rows.Scan(&studentID, &response.Semester)
		if err != nil {
			return nil, err
		}

		student, err := GetStudent(studentID)
		if err != nil {
			return nil, err
		}

		response.StudentName = fmt.Sprintf("%s %s", student.FName, student.LName)
		response.StudentID = student.ID

		courses = append(courses, response)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return courses, nil
}

func GetCourseStudentNamesEnrollments(courseId int16, semester string) ([]string, error) {
	query := `
		SELECT student_id FROM enrollment
		WHERE course_id = ? AND semester = ?
	`
	rows, err := DB.Query(query, courseId, semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentNames []string
	for rows.Next() {
		var studentID int16
		err := rows.Scan(&studentID)
		if err != nil {
			return nil, err
		}

		student, err := GetStudent(studentID)
		if err != nil {
			return nil, err
		}

		studentNames = append(studentNames, fmt.Sprintf("%s %s", student.FName, student.LName))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return studentNames, nil
}

type CourseEnrollmentCount struct {
	CourseID      int
	CourseName    string
	TotalStudents int
}

func GetStudentCountCID() ([]CourseEnrollmentCount, error) {
	query := `
	SELECT course_id, COUNT(student_id) AS total_students
	FROM enrollment
	GROUP BY course_id;
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courseEnrollmentsCount []CourseEnrollmentCount
	for rows.Next() {
		var ce CourseEnrollmentCount
		err := rows.Scan(&ce.CourseID, &ce.TotalStudents)
		if err != nil {
			return nil, err
		}
		course, err := GetCourse(int16(ce.CourseID))
		if err != nil {
			return nil, err
		}
		ce.CourseName = course.CourseName
		courseEnrollmentsCount = append(courseEnrollmentsCount, ce)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courseEnrollmentsCount, nil
}

type StudentSemesterCount struct {
	Semester     string
	TotalStudent int16
}

func GetStudentCountSemester() ([]StudentSemesterCount, error) {
	query := `
		SELECT semester, COUNT(student_id) AS total_students
		FROM enrollment
		GROUP BY semester
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []StudentSemesterCount
	for rows.Next() {
		var s StudentSemesterCount
		if err = rows.Scan(&s.Semester, &s.TotalStudent); err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

type CourseGrading struct {
	CourseName  string
	Score       float32
	CourseGrade string
}

func CourseGrades() ([]CourseGrading, error) {
	query := `
		SELECT course_id, AVG(score) AS average_score
		FROM enrollment
		GROUP BY course_id
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []CourseGrading
	for rows.Next() {
		var s CourseGrading
		var courseID int16
		if err = rows.Scan(&courseID, &s.Score); err != nil {
			return nil, err
		}

		s.CourseGrade = SelectGrade(s.Score)
		course, err := GetCourse(courseID)
		if err != nil {
			return nil, err
		}
		s.CourseName = course.CourseName

		result = append(result, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

type StudentsGrading struct {
	StudentName   string
	StudentScore  float32
	StudentsGrade string
}

func StudentsGrades() ([]StudentsGrading, error) {
	query := `
		SELECT student_id, AVG(score) AS average_score
		FROM enrollment
		GROUP BY student_id
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []StudentsGrading
	for rows.Next() {
		var s StudentsGrading
		var studentID int16
		if err = rows.Scan(&studentID, &s.StudentScore); err != nil {
			return nil, err
		}

		s.StudentsGrade = SelectGrade(s.StudentScore)
		student, err := GetStudent(studentID)
		if err != nil {
			return nil, err
		}
		s.StudentName = fmt.Sprintf("%s %s", student.FName, student.LName)

		result = append(result, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
