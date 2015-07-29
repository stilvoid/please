package main

import (
	"fmt"
	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/formatter"
	"github.com/stilvoid/please/parser"
	"io/ioutil"
	"os"
)

func init() {
	commands["parse"] = parseCommand
}

func parseHelp() {
	fmt.Println("Usage: please parse [-i <INPUT FORMAT>] [-o <OUTPUT FORMAT>] [path.to.extract]")
	fmt.Println()
	fmt.Println("If omitted, the input type defaults to \"auto\". The output type defaults to \"bash\".")
	fmt.Println()
	fmt.Println("Input types:")
	for _, format := range parser.Names() {
		fmt.Printf("    %s\n", format)
	}
	fmt.Println()
	fmt.Println("Output types:")
	for _, format := range formatter.Names() {
		fmt.Printf("    %s\n", format)
	}
}

func parseCommand(args []string) {
	// Flags
	inFormat := getopt.String('i', "auto")
	outFormat := getopt.String('o', "bash")

	opts := getopt.CommandLine

	opts.SetUsage(parseHelp)

	opts.Parse(args)

	if termutil.Isatty(os.Stdin.Fd()) {
		getopt.Usage()
		os.Exit(1)
	}

	var err error

	// Read from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	// Try parsing
	parsed, err := parser.Parse(input, *inFormat)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Path
	if getopt.NArgs() > 0 {
		parsed, err = parser.Filter(parsed, getopt.Arg(0))

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// ...and format back out :)
	output, err := formatter.Format(parsed, *outFormat)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}
