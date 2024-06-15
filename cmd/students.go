/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// studentsCmd represents the students command
var studentsCmd = &cobra.Command{
	Use:   "students",
	Short: "Calculates the students grades",
	Long:  `Calculates the students grades`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := db.StudentsGrades()
		if err != nil {
			fmt.Println("coulnt calculate: ", err)
			return
		}

		for _, data := range results {
			fmt.Printf("%v - %v - %v\n", data.StudentName, data.StudentsGrade, data.StudentScore)
		}
	},
}

func init() {
	gradeCmd.AddCommand(studentsCmd)
}
