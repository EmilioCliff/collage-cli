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

// addCourseCmd represents the addCourse command
var addCourseCmd = &cobra.Command{
	Use:   "addCourse",
	Short: "Add course to database",
	Long:  `Add course to database`,
	Run: func(cmd *cobra.Command, args []string) {
		course, err := addCourse()
		if err != nil {
			fmt.Println("total credit should be in")
		}

		_, err = db.AddCourse(course)
		if err != nil {
			fmt.Println("Error adding course to database: ", err)
			return
		}
	},
}

func addCourse() (db.Course, error) {
	courseNamePrompt := PromptContent{
		"Please provide course name",
		"Course Name: ",
	}
	courseName := promptGetInput(courseNamePrompt)

	totalCreditPrompt := PromptContent{
		"Please provide course total credit",
		"Course Credit (ie. 29000): ",
	}
	totalCredit := promptGetInput(totalCreditPrompt)
	totalCreditint, err := strconv.Atoi(totalCredit)
	if err != nil {
		return db.Course{}, err
	}

	course := db.Course{
		CourseName:  courseName,
		TotalCredit: int32(totalCreditint),
	}

	return course, nil
}
