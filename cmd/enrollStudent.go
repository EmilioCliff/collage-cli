/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// enrollStudentCmd represents the enrollStudent command
var enrollStudentCmd = &cobra.Command{
	Use:   "enrollStudent",
	Short: "Enroll Student",
	Long:  `Enroll Student`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enrollStudent called")
		enroll, err := enrollStudent()
		if err != nil {
			fmt.Println("Please enter correct course/student id: HINT<should be numbers>", err)
			return
		}
		_, err = db.EnrollStudent(enroll)
		if err != nil {
			fmt.Println("Error enrolling student: ", err)
			return
		}
	},
}

func enrollStudent() (db.Enroll, error) {
	courseIDPrompt := PromptContent{
		"Please provide course id",
		"Course ID: ",
	}
	courseID := promptGetInput(courseIDPrompt)
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		return db.Enroll{}, err
	}

	studentIDPrompt := PromptContent{
		"Please provide students id",
		"Student ID: ",
	}
	studentID := promptGetInput(studentIDPrompt)
	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		return db.Enroll{}, err
	}

	categoryPromptContent := PromptContent{
		"Add Semester Enrolled",
		"Semester: ",
	}
	semester := promptSelect(categoryPromptContent, []string{"fall", "winter"})

	// semesterPrompt := PromptContent{
	// 	"Please provide semeseter",
	// 	"Semester (fall, winter): ",
	// }
	// semester := promptGetInput(semesterPrompt)

	scorePrompt := PromptContent{
		"Please provide student score",
		"Student Score: ",
	}
	score := promptGetInput(scorePrompt)
	scoreFloat, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return db.Enroll{}, err
	}

	enroll := db.Enroll{
		CourseID:  int16(courseIDInt),
		StudentID: int16(studentIDInt),
		Semester:  semester,
		Score:     float32(scoreFloat),
	}

	return enroll, nil
}
