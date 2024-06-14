/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// listEnrollmentsCmd represents the listEnrollments command
var listEnrollmentsCmd = &cobra.Command{
	Use:   "listEnrollments",
	Short: "List Enrollments",
	Long:  `List Enrollments`,
	Run: func(cmd *cobra.Command, args []string) {
		enrollments, err := db.ListEnrollments()
		if err != nil {
			fmt.Println("Could not list enrollments: ", err)
			return
		}
		fmt.Println("course_id, student_id, semester, score")
		for i := 0; i < len(enrollments); i++ {
			fmt.Printf("%v, %v, %v, %v\n", enrollments[i].CourseID, enrollments[i].StudentID, enrollments[i].Semester, enrollments[i].Score)
		}
	},
}
