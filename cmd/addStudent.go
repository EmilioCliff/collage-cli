/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// addStudentCmd represents the addStudent command
var addStudentCmd = &cobra.Command{
	Use:   "addStudent",
	Short: "Add student to database",
	Long:  `Add student to database`,
	Run: func(cmd *cobra.Command, args []string) {
		student := addStudent()
		_, err := db.AddStudent(student)
		if err != nil {
			fmt.Println("Error adding student to db: ", err)
			return
		}
	},
}

func addStudent() db.Student {
	fNamePrompt := PromptContent{
		"Please provide students first name",
		"First Name: ",
	}
	fName := promptGetInput(fNamePrompt)

	lNamePrompt := PromptContent{
		"Please provide students last name",
		"Last Name: ",
	}
	lName := promptGetInput(lNamePrompt)

	dobPrompt := PromptContent{
		"Please provide students Date Of Birth",
		"Date Of Birth (2006-10-04): ",
	}
	dob := promptGetInput(dobPrompt)

	student := db.Student{
		FName: fName,
		LName: lName,
		DOB:   dob,
	}

	return student
}
