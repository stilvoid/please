package cmd

import (
	"fmt"
	"io/ioutil"
	"maps"
	"os"
	"slices"
	"sort"

	"github.com/andrew-d/go-termutil"
	"github.com/jmespath/go-jmespath"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please"
)

var parserNames []string
var formatNames []string

func init() {
	Commands["parse"] = parseCommand

	parserNames = slices.Collect(maps.Keys(please.Parsers))
	sort.Strings(parserNames)

	formatNames = slices.Collect(maps.Keys(please.Parsers))
	sort.Strings(formatNames)
}

func parseHelp() {
	fmt.Println("Usage: please parse [-i <INPUT FORMAT>] [-o <OUTPUT FORMAT>] [PATH]")
	fmt.Println()
	fmt.Println("If INPUT TYPE is omitted, it defaults to \"auto\".")
	fmt.Println()
	fmt.Println("If OUTPUT TYPE is omitted, it will default to the same as the input type - effectively acting as a pretty-printer.")
	fmt.Println()
	fmt.Println("PATH, if provided, is a JMESPath query, e.g. orders.*.id")
	fmt.Println()
	fmt.Println("Available input types:")
	fmt.Printf("    auto\n")
	for _, format := range parserNames {
		fmt.Printf("    %s\n", format)
	}
	fmt.Println()
	fmt.Println("Available output types:")
	for _, format := range formatNames {
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
	var parsed any

	// Deal with format detection
	if *inFormat == "auto" {
		*inFormat, parsed, err = identify(input)
	} else {
		// Try parsing
		parser, ok := please.Parsers[*inFormat]
		if !ok {
			fmt.Fprintf(os.Stderr, "No such parser: %s\n", *inFormat)
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
		parsed, err = jmespath.Search(getopt.Arg(0), parsed)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *outFormat == "auto" {
		*outFormat = *inFormat
	}

	// ...and format back out :)
	formatter, ok := please.Formatters[*outFormat]
	if !ok {
		fmt.Fprintf(os.Stderr, "No such formatter: %s\n", *outFormat)
		os.Exit(1)
	}

	output, err := formatter(parsed)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}
