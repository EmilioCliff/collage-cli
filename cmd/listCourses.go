/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// listCoursesCmd represents the listCourses command
var listCoursesCmd = &cobra.Command{
	Use:   "listCourses",
	Short: "List Courses",
	Long:  `List Courses`,
	Run: func(cmd *cobra.Command, args []string) {
		courses, err := db.ListCourses()
		if err != nil {
			fmt.Println("Could not list courses: ", err)
			return
		}
		fmt.Println("course_id, course_name, total_credit")
		for i := 0; i < len(courses); i++ {
			fmt.Printf("%v, %v, %v\n", courses[i].ID, courses[i].CourseName, courses[i].TotalCredit)
		}
	},
}
