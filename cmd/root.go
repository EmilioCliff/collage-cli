/*
Copyright © 2024 Emilio Cliff emiliocliff@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "A CLI application for managing students and courses",
	Long:  `Welcome to the Student Management CLI. Use this application to manage students and their courses.`,
}

func InitializeCommands() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(addCourseCmd)
	RootCmd.AddCommand(addStudentCmd)
	RootCmd.AddCommand(enrollStudentCmd)
	RootCmd.AddCommand(listCoursesCmd)
	RootCmd.AddCommand(listEnrollmentsCmd)
	RootCmd.AddCommand(listStudentsCmd)
	RootCmd.AddCommand(searchCourseCmd)
	RootCmd.AddCommand(searchStudentCmd)
	RootCmd.AddCommand(countCmd)
	RootCmd.AddCommand(gradeCmd)
	RootCmd.AddCommand(clearCmd)
}
