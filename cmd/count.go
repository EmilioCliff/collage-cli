/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "This commands deals with counting",
	Long:  `This commands deals with counting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("count called")
	},
}
