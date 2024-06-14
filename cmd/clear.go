/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var OScmd *exec.Cmd
		if runtime.GOOS == "windows" {
			OScmd = exec.Command("cmd", "/c", "cls")
		} else {
			OScmd = exec.Command("clear")
		}
		OScmd.Stdout = os.Stdout
		OScmd.Run()
	},
}
