/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// coursesCmd represents the courses command
var coursesCmd = &cobra.Command{
	Use:   "courses",
	Short: "Calculates the grades of each course",
	Long:  `Calculates the grades of each course`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := db.CourseGrades()
		if err != nil {
			fmt.Println("coulnt calculate: ", err)
			return
		}

		for _, data := range results {
			fmt.Printf("%v - %v - %v\n", data.CourseName, data.CourseGrade, data.Score)
		}
	},
}

func init() {
	gradeCmd.AddCommand(coursesCmd)
}
