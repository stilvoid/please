package main

import (
	"fmt"
	"os"

	"github.com/stilvoid/please/cmd"
)

func printHelp() {
	fmt.Println("Usage: please <COMMAND> [arg...]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("    identify   Identify the structured data on stdin")
	fmt.Println("    request    Make a web request and output the result")
	fmt.Println("    respond    Listen for a web request and respond to it")
	fmt.Println("    parse      Get values from structured data and convert between formats")
	fmt.Println()
	fmt.Println("Run 'please COMMAND' for more information on a command.")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	args := os.Args[1:]

	// Deal with aliases
	if alias, ok := cmd.Aliases[command]; ok {
		newArgs := make([]string, 1, len(args)+1)
		newArgs[0] = alias

		args = append(newArgs, args...)

		command = alias
	}

	commandFunc, ok := cmd.Commands[command]

	if !ok {
		printHelp()
		os.Exit(1)
	}

	commandFunc(args)
}
