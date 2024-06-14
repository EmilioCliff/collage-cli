/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// searchCourseCmd represents the searchCourse command
var searchCourseCmd = &cobra.Command{
	Use:   "searchCourse",
	Short: "Searches for a course in the database",
	Long:  `Searches for a course in the database`,
}
