package main

import (
	"fmt"
	"github.com/stilvoid/please/cmd"
	"os"
)

var commands map[string]func([]string)

func init() {
	commands = map[string]func([]string){
		"request": cmd.Request,
		"respond": cmd.Respond,
		"parse":   cmd.Parse,
	}
}

func printHelp() {
	fmt.Println("Usage: please <COMMAND> [arg...]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("    request    Make a web request")
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

	if alias, ok := cmd.RequestAliases[command]; ok {
		new_args := make([]string, 1)
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
