/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/EmilioCliff/collage-cli/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the db",
	Long:  `Creates tables`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.OpenDatabas()
		if err != nil {
			log.Println(err)
		}
		db.CreateTables()
	},
}
