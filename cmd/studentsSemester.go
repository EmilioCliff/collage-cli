/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// studentsSemesterCmd represents the studentsSemester command
var studentsSemesterCmd = &cobra.Command{
	Use:   "studentsSemester",
	Short: "This will display the total number of students for each semester",
	Long:  `This will display the total number of students for each semester`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("studentsSemester called")
	},
}

func init() {
	countCmd.AddCommand(studentsSemesterCmd)
}
