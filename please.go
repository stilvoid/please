package main

import (
	"fmt"
	"github.com/stilvoid/please/cmd"
	"os"
)

var commands map[string]func([]string)
var aliases map[string]string

func init() {
	commands = map[string]func([]string){
		"request": cmd.Request,
		"respond": cmd.Respond,
		"parse":   cmd.Parse,
	}

	aliases = map[string]string{
		"get":    "request",
		"post":   "request",
		"put":    "request",
		"delete": "request",
	}
}

func printHelp() {
	fmt.Println("Usage: please <command> [options...]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("    go request")
	fmt.Println("    go respond")
	fmt.Println("    go parse")
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	args := os.Args[1:]

	if alias, ok := aliases[command]; ok {
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
