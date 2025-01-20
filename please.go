package please

import (
	"fmt"
	"maps"
	"slices"

	"github.com/stilvoid/please/format"
	"github.com/stilvoid/please/parse"
)

var formatters = map[string]func(any) (string, error){
	"bash":  format.Bash,
	"dot":   format.Dot,
	"json":  format.Json,
	"xml":   format.Xml,
	"yaml":  format.Yaml,
	"query": format.Query,
	"toml":  format.Toml,
}

var parsers = map[string]func([]byte) (any, error){
	"csv":   parse.Csv,
	"html":  parse.Html,
	"json":  parse.Json,
	"mime":  parse.Mime,
	"xml":   parse.Xml,
	"yaml":  parse.Yaml,
	"query": parse.Query,
	"toml":  parse.Toml,
}

// Formatters are the names of data formats that please can output
var Formatters []string

// Parsers are the names of data formats that please can read
var Parsers []string

func init() {
	Formatters = slices.Collect(maps.Keys(formatters))
	Parsers = slices.Collect(maps.Keys(parsers))
}

// Format converts input data into format as a byte array
func Format(format string, input any) (string, error) {
	formatter, ok := formatters[format]
	if !ok {
		return "", fmt.Errorf("Invalid format: %s", format)
	}

	return formatter(input)
}

// Parse reads input as format and returns data
func Parse(format string, input []byte) (any, error) {
	parser, ok := parsers[format]
	if !ok {
		return nil, fmt.Errorf("Invalid format: %s", format)
	}

	return parser(input)
}
