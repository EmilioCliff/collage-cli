/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// cidSemesterCmd represents the cidSemester command
var cidSemesterCmd = &cobra.Command{
	Use:   "cidSemester",
	Short: "searches for students taking a course for a certain semester",
	Long:  `searches for students taking a course for a certain semester`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		searchCourse, err := db.SearchCourseByIDandSemester(int16(id), args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(searchCourse)
	},
}

func init() {
	searchCourseCmd.AddCommand(cidSemesterCmd)
}
