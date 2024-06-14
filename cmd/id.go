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

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:   "id",
	Short: "search using student's ID",
	Long:  `uses the students id to search`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		searchResult, err := db.SearchStudentByID(int16(id))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(searchResult)
	},
}

func init() {
	searchStudentCmd.AddCommand(idCmd)
}
