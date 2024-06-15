/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// studentsCourseCmd represents the studentsCourse command
var studentsCourseCmd = &cobra.Command{
	Use:   "studentsCourse",
	Short: "A brief description of your command",
	Long:  `This will display the total number of students for each cid...courseID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("studentsCourse called")
	},
}

func init() {
	countCmd.AddCommand(studentsCourseCmd)
}
