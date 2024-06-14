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

var cidCmd = &cobra.Command{
	Use:   "cid",
	Short: "searches the students enrolled in a certain course",
	Long:  `searches the students enrolled in a certain course`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		searchCourse, err := db.SearchCourseByID(int16(id))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(searchCourse)
	},
}

func init() {
	searchCourseCmd.AddCommand(cidCmd)
}

func createNewNote() {
	wordPromptContent := PromptContent{
		"Please provide a word",
		"What word would you like to make a note of",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := PromptContent{
		"Please provide a definition",
		fmt.Sprintf("What is the definition of %s", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := PromptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belong to", word),
	}
	category := promptSelect(categoryPromptContent, []string{"action", "animal", "toy"})

	fmt.Print("Created new note successfully: \n", word, definition, category)
}
