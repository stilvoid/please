package cmd

import (
	"fmt"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/formatter"
	"github.com/stilvoid/please/parser"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
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

func parseAuto(input []byte) (interface{}, string, error) {
	var parsed interface{}
	var err error

	for _, name := range parser_preference {
		parsed, err = parsers[name](input)

		if err == nil {
			return parsed, name, err
		}
	}

	return nil, "", fmt.Errorf("Input format could not be identified")
}

func forceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		new_map := make(map[string]interface{}, val.Len())

		for _, key := range val.MapKeys() {
			string_key := fmt.Sprint(key.Interface())
			new_map[string_key] = forceStringKeys(val.MapIndex(key).Interface())
		}

		return new_map
	case reflect.Array, reflect.Slice:
		new_slice := make([]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			new_slice[i] = forceStringKeys(value)
		}

		return new_slice
	default:
		return in
	}
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

func sortKeys(in interface{}) []string {
	v := reflect.ValueOf(in)

	if v.Kind() != reflect.Map {
		panic("Not a map!")
	}

	vkeys := v.MapKeys()

	keys := make([]string, 0, len(vkeys))

	for _, key := range vkeys {
		keys = append(keys, key.String())
	}

	sort.Strings(keys)

	return keys
}

func parse(input []byte, format string) (interface{}, string, error) {
	if format == "auto" {
		return parseAuto(input)
	}

	output, err := parsers[format](input)

	return output, format, err
}

func format(input interface{}, format string) string {
	return formatters[format](input)
}

func printParsers() {
	fmt.Printf(" Input types: %s\n", strings.Join(sortKeys(parsers), ", "))
}

func printFormatters() {
	fmt.Printf(" Output types: %s\n", strings.Join(sortKeys(formatters), ", "))
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
		parsed = filter(parsed, getopt.Arg(0))
	}

	// Pretty much everything hates non-string keys :S
	if detected_in_format == "yaml" && *out_format != "yaml" {
		parsed = forceStringKeys(parsed)
	}

	// ...and format back out :)
	fmt.Println(format(parsed, *out_format))
}
