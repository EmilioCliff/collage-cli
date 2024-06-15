/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// studentsSemesterCmd represents the studentsSemester command
var studentsSemesterCmd = &cobra.Command{
	Use:   "studentsSemester",
	Short: "This will display the total number of students for each semester",
	Long:  `This will display the total number of students for each semester`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := db.GetStudentCountSemester()
		if err != nil {
			fmt.Println("Could not calculate semester: ", err)
			return
		}

		for _, data := range result {
			fmt.Printf("%v - %v\n", data.Semester, data.TotalStudent)
		}
	},
}

func init() {
	countCmd.AddCommand(studentsSemesterCmd)
}
