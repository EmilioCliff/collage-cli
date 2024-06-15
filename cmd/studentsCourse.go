/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// studentsCourseCmd represents the studentsCourse command
var studentsCourseCmd = &cobra.Command{
	Use:   "studentsCourse",
	Short: "A brief description of your command",
	Long:  `This will display the total number of students for each cid...courseID`,
	Run: func(cmd *cobra.Command, args []string) {
		totalStudents, err := db.GetStudentCountCID()
		if err != nil {
			fmt.Println("failed to calculate: ", err)
		}
		for _, data := range totalStudents {
			fmt.Printf("%v - %v\n", data.CourseName, data.TotalStudents)
		}
	},
}

func init() {
	countCmd.AddCommand(studentsCourseCmd)
}
