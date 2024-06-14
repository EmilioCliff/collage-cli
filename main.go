/*
Copyright © 2024 Emilio Cliff emiliocliff@gmail.com
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmd "github.com/EmilioCliff/collage-cli/cmd"
)

const promptColor = "\033[32m"
const resetColor = "\033[0m"

func main() {
	cmd.InitializeCommands()

	fmt.Println("Welcome to the Student Management CLI")
	fmt.Println("Type 'help' to view program commands")
	fmt.Println("Type 'exit' to quit the application")

	initializeProgram()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(string(promptColor) + "collage-cli: " + string(resetColor))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		args := strings.Split(input, " ")
		cmd.RootCmd.SetArgs(args)

		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Exiting the Student Management CLI. Goodbye!")
}

func initializeProgram() {
	cmd.RootCmd.SetArgs([]string{"init"})
	if err := cmd.RootCmd.Args; err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
