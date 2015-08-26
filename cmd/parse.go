package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/formatters"
	"github.com/stilvoid/please/parsers"
	"github.com/stilvoid/please/util"
)

func init() {
	Commands["parse"] = parseCommand
}

func parseHelp() {
	fmt.Println("Usage: please parse [-i <INPUT FORMAT>] [-o <OUTPUT FORMAT>] [PATH]")
	fmt.Println()
	fmt.Println("If INPUT TYPE is omitted, it defaults to \"auto\".")
	fmt.Println()
	fmt.Println("If OUTPUT TYPE is omitted, it will default to the same as the input type - effectively acting as a pretty-printer.")
	fmt.Println()
	fmt.Println("If PATH is provided, it should be given in dot-notation, e.g. orders.*.id")
	fmt.Println()
	fmt.Println("Available input types:")
	fmt.Printf("    auto\n")
	for _, format := range parsers.Names() {
		fmt.Printf("    %s\n", format)
	}
	fmt.Println()
	fmt.Println("Available output types:")
	for _, format := range formatters.Names() {
		fmt.Printf("    %s\n", format)
	}
}

func parseCommand(args []string) {
	// Flags
	inFormat := getopt.String('i', "auto")
	outFormat := getopt.String('o', "auto")

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
	var parsed interface{}

	// Deal with format detection
	if *inFormat == "auto" {
		*inFormat, parsed, err = parsers.Identify(input)
	} else {
		// Try parsing
		parser, innerErr := parsers.Get(*inFormat)

		if innerErr != nil {
			fmt.Fprintln(os.Stderr, innerErr)
			os.Exit(1)
		}

		parsed, err = parser(input)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Path
	if getopt.NArgs() > 0 {
		parsed, err = util.Filter(parsed, getopt.Arg(0))

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *outFormat == "auto" {
		*outFormat = *inFormat
	}

	// ...and format back out :)
	formatter, err := formatters.Get(*outFormat)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output, err := formatter(parsed)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}
