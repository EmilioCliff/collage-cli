/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// studentsCmd represents the students command
var studentsCmd = &cobra.Command{
	Use:   "students",
	Short: "Counts the total number of students",
	Long:  `Counts the total number of students`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("students called")
	},
}

func init() {
	countCmd.AddCommand(studentsCmd)
}
