/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// gradeCmd represents the grade command
var gradeCmd = &cobra.Command{
	Use:   "grade",
	Short: "A brief description of your command",
	Long:  `This calculates the grades of the students or courses`,
}
