package cmd

import (
	"code.google.com/p/getopt"
	"fmt"
	"github.com/stilvoid/please/formatter"
	"github.com/stilvoid/please/parser"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Parser func([]byte) (interface{}, error)
type Formatter func(interface{}) string

var parsers map[string]Parser
var parser_preference []string
var formatters map[string]Formatter

func parseAuto(input []byte) (interface{}, error) {
	var parsed interface{}
	var err error

	for _, name := range parser_preference {
		parsed, err = parsers[name](input)

		if err == nil {
			break
		}
	}

	return parsed, err
}

func filter(in interface{}, path string) interface{} {
	if path == "" {
		return in
	}

	split_path := strings.SplitN(path, ".", 2)

	this_path := split_path[0]
	var next_path string

	if len(split_path) > 1 {
		next_path = split_path[1]
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		vv := in.(map[string]interface{})

		next, ok := vv[this_path]

		if !ok {
			break
		}

		return filter(next, next_path)
	case reflect.Array, reflect.Slice:
		index, err := strconv.Atoi(this_path)

		if err != nil || index < 0 || index >= val.Len() {
			break
		}

		return filter(val.Index(index).Interface(), next_path)
	}

	fmt.Fprintf(os.Stderr, "Key does not exist %s\n", this_path)
	os.Exit(1)

	return nil
}

func init() {
	parsers = map[string]Parser{
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

	formatters = map[string]Formatter{
		"bash": formatter.Bash,
		"dot":  formatter.Dot,
		"xml":  formatter.Xml,
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

	// Try parsing
	parsed, err := parsers[*in_format](input)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Input could not be parsed")
		fmt.Println(err)
		os.Exit(1)
	}

	// Path
	if getopt.NArgs() > 0 {
		parsed = filter(parsed, getopt.Arg(0))
	}

	// ...and format back out :)
	fmt.Println(formatters[*out_format](parsed))
}
