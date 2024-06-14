/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// lnameCmd represents the lname command
var lnameCmd = &cobra.Command{
	Use:   "lname",
	Short: "search using student's ID",
	Long:  `search using student's ID`,
	Run: func(cmd *cobra.Command, args []string) {
		searchResult, err := db.SearchStudentByLName(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < len(searchResult); i++ {
			fmt.Println(searchResult[i])
		}
	},
}

func init() {
	searchStudentCmd.AddCommand(lnameCmd)
}
