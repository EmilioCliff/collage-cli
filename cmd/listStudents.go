/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

var listStudentsCmd = &cobra.Command{
	Use:   "listStudents",
	Short: "List Students",
	Long:  `List Students`,
	Run: func(cmd *cobra.Command, args []string) {
		students, err := db.ListStudents()
		if err != nil {
			fmt.Println("Could not list students: ", err)
			return
		}
		fmt.Println("id,first-name,last-name,date-of-birth")
		for i := 0; i < len(students); i++ {
			fmt.Printf("%v, %v, %v, %v\n", students[i].ID, students[i].FName, students[i].LName, students[i].DOB)
		}
	},
}
