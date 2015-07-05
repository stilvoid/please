package cmd

import (
	"fmt"
	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/formatter"
	"github.com/stilvoid/please/parser"
	"github.com/stilvoid/please/util"
	"io/ioutil"
	"os"
	"strings"
)

func Parse(args []string) {
	// Flags
	in_format := getopt.String('i', "auto", "Parse the input as 'types'", "type")
	out_format := getopt.String('o', "bash", "Use 'type' as the output format", "type")

	opts := getopt.CommandLine

	opts.SetUsage(func() {
		getopt.CommandLine.PrintUsage(os.Stderr)

		fmt.Println()

		fmt.Printf(" Input types: %s\n", strings.Join(util.SortKeys(parser.Parsers), ", "))

		fmt.Printf(" Output types: %s\n", strings.Join(util.SortKeys(formatter.Formatters), ", "))
	})

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
	parsed, _, err := parser.Parse(input, *in_format)

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

	// ...and format back out :)
	output, err := formatter.Format(parsed, *out_format)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}
