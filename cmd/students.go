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
	Short: "Counts the total number of students",
	Long:  `Counts the total number of students`,
	Run: func(cmd *cobra.Command, args []string) {
		totalStudents, err := db.CountStudents()
		if err != nil {
			fmt.Println("failed to count students: ", err)
			return
		}
		fmt.Println(totalStudents)
	},
}

func init() {
	countCmd.AddCommand(studentsCmd)
}
