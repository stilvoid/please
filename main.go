package main

import (
	"fmt"
	"os"
)

type command func([]string)

var commands = make(map[string]command)

var aliases = make(map[string]string)

func printHelp() {
	fmt.Println("Usage: please <COMMAND> [arg...]")
	fmt.Println()
	fmt.Println("Commands:")
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
	if alias, ok := aliases[command]; ok {
		new_args := make([]string, 1, len(args)+1)
		new_args[0] = alias

		args = append(new_args, args...)

		command = alias
	}

	command_func, ok := commands[command]

	if !ok {
		printHelp()
		os.Exit(1)
	}

	command_func(args)
}
