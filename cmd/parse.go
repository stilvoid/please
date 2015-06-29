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

type Parser func([]byte) (interface{}, error)
type Formatter func(interface{}) string

var parsers map[string]Parser
var parser_preference []string
var formatters map[string]Formatter

func init() {
	parsers = map[string]Parser{
		"csv":  parser.Csv,
		"html": parser.Html,
		"json": parser.Json,
		"mime": parser.Mime,
		"xml":  parser.Xml,
		"yaml": parser.Yaml,
	}

	parser_preference = []string{
		"json",
		"xml",
		"yaml",
		"csv",
		"html",
		"mime",
	}

	formatters = map[string]Formatter{
		"bash": formatter.Bash,
		"dot":  formatter.Dot,
		"json": formatter.Json,
		"xml":  formatter.Xml,
		"yaml": formatter.Yaml,
	}
}

func parse(input []byte, format string) (interface{}, string, error) {
	if format == "auto" {
		for _, name := range parser_preference {
			parsed, err := parsers[name](input)

			if err == nil {
				return parsed, name, err
			}
		}

		return nil, "", fmt.Errorf("Input format could not be identified")
	}

	output, err := parsers[format](input)
	return output, format, err
}

func format(input interface{}, format string) string {
	return formatters[format](input)
}

func printParsers() {
	fmt.Printf(" Input types: %s\n", strings.Join(util.SortKeys(parsers), ", "))
}

func printFormatters() {
	fmt.Printf(" Output types: %s\n", strings.Join(util.SortKeys(formatters), ", "))
}

func Parse(args []string) {
	// Flags
	in_format := getopt.String('i', "auto", "Parse the input as 'types'", "type")
	out_format := getopt.String('o', "bash", "Use 'type' as the output format", "type")

	opts := getopt.CommandLine

	opts.SetUsage(func() {
		getopt.CommandLine.PrintUsage(os.Stderr)
		fmt.Println()
		printParsers()
		printFormatters()
	})

	opts.Parse(args)

	if termutil.Isatty(os.Stdin.Fd()) {
		getopt.Usage()
		os.Exit(1)
	}

	// Validate parser
	if _, ok := parsers[*in_format]; !ok && *in_format != "auto" {
		fmt.Printf("Unknown input format: %s\n\n", *in_format)
		printParsers()
		os.Exit(1)
	}

	// Validate formatter
	if _, ok := formatters[*out_format]; !ok {
		fmt.Printf("Unknown output format: %s\n\n", *out_format)
		printFormatters()
		os.Exit(1)
	}

	var err error

	// Read from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}

	// Try parsing
	parsed, detected_in_format, err := parse(input, *in_format)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Input could not be parsed")
		fmt.Println(err)
		os.Exit(1)
	}

	// Path
	if getopt.NArgs() > 0 {
		parsed = util.Filter(parsed, getopt.Arg(0))
	}

	// Pretty much everything hates non-string keys :S
	if detected_in_format == "yaml" && *out_format != "yaml" {
		parsed = util.ForceStringKeys(parsed)
	}

	// ...and format back out :)
	fmt.Println(format(parsed, *out_format))
}
