package cmd

import (
	"code.google.com/p/getopt"
	"fmt"
	"github.com/stilvoid/please/formatter"
	"github.com/stilvoid/please/parser"
	"io/ioutil"
	"os"
)

var parsers map[string]func([]byte, string) (interface{}, error)
var parser_preference []string
var formatters map[string]func(interface{}, string) string

func parseAuto(input []byte, path string) (interface{}, error) {
	var parsed interface{}
	var err error

	for _, name := range parser_preference {
		parsed, err = parsers[name](input, path)

		if err == nil {
			break
		}
	}

	return parsed, err
}

func init() {
	parsers = map[string]func([]byte, string) (interface{}, error){
		"auto": parseAuto,
		"json": parser.Json,
		"xml":  parser.Xml,
		"csv":  parser.Csv,
		"html": parser.Html,
		"mime": parser.Mime,
	}

	parser_preference = []string{
		"json",
		"xml",
		"csv",
		"html",
		"mime",
	}

	formatters = map[string]func(interface{}, string) string{
		"bash": formatter.Bash,
		"yaml": formatter.Yaml,
	}
}

func Parse(args []string) {
	// Flags
	in_format := getopt.String('i', "auto", "Parse the input as 'types'", "type")
	out_format := getopt.String('o', "bash", "Use 'type' as the output format", "type")

	opts := getopt.CommandLine

	opts.Parse(args)

	// Validate parser
	if _, ok := parsers[*in_format]; !ok {
		fmt.Printf("Unknown input format: %s\n", *in_format)
		os.Exit(1)
	}

	// Validate formatter
	if _, ok := formatters[*out_format]; !ok {
		fmt.Printf("Unknown output format: %s\n", *out_format)
		os.Exit(1)
	}

	var err error

	// Read from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}

	// Path
	var path string

	if getopt.NArgs() > 0 {
		path = getopt.Arg(0)
	}

	// Try parsing
	parsed, err := parsers[*in_format](input, path)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Input could not be parsed")
		fmt.Println(err)
		os.Exit(1)
	}

	// ...and format back out :)
	fmt.Println(formatters[*out_format](parsed, path))
}
