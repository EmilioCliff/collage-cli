package db

type StudentResponse struct {
	StudentID   int16
	StudentName string
	Semester    string
}

type CourseSearch struct {
	CourseName     string
	StudentsEnroll []StudentResponse
	Credit         int32
	Grade          string
}

type CourseSearchWithSemester struct {
	CourseName  string
	StudentName []string
	Credit      int32
	Grade       string
}

func AddCourse(arg Course) (Course, error) {
	query := `
		INSERT INTO courses
		(course_name, total_credit)
		VALUES (?, ?)
		RETURNING id, course_name, total_credit
	`
	row := DB.QueryRow(query, arg.CourseName, arg.TotalCredit)

	var result Course
	if err := row.Scan(&result.ID, &result.CourseName, &result.TotalCredit); err != nil {
		return Course{}, err
	}

	return result, nil
}

func ListCourses() ([]Course, error) {
	query := `
		SELECT * FROM courses 
		ORDER BY course_name
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Course
	for rows.Next() {
		var course Course
		if err = rows.Scan(&course.ID, &course.CourseName, &course.TotalCredit); err != nil {
			return nil, err
		}

		result = append(result, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetCourse(courseID int16) (Course, error) {
	query := `
		SELECT * FROM courses
		WHERE id = ?
	`
	row := DB.QueryRow(query, courseID)

	var result Course
	if err := row.Scan(&result.ID, &result.CourseName, &result.TotalCredit); err != nil {
		return Course{}, err
	}

	return result, nil
}

func SearchCourseByID(courseID int16) (CourseSearch, error) {
	var result CourseSearch
	course, err := GetCourse(courseID)
	if err != nil {
		return CourseSearch{}, err
	}

	result.CourseName = course.CourseName
	result.Credit = course.TotalCredit
	score, err := GetCourseAverageScore(courseID)
	if err != nil {
		return CourseSearch{}, err
	}

	result.Grade = SelectGrade(score)
	result.StudentsEnroll, err = GetCourseEnrollments(courseID)
	if err != nil {
		return CourseSearch{}, err
	}

	return result, nil
}

func SearchCourseByIDandSemester(courseID int16, semester string) (CourseSearchWithSemester, error) {
	var result CourseSearchWithSemester
	course, err := GetCourse(courseID)
	if err != nil {
		return CourseSearchWithSemester{}, err
	}

	result.CourseName = course.CourseName
	result.Credit = course.TotalCredit
	score, err := GetCourseAndSemesterAverageScore(courseID, semester)
	if err != nil {
		return CourseSearchWithSemester{}, err
	}

	result.Grade = SelectGrade(score)
	result.StudentName, err = GetCourseStudentNamesEnrollments(courseID, semester)
	if err != nil {
		return CourseSearchWithSemester{}, err
	}

	return result, nil
}
