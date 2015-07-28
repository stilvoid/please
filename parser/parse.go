// Package parser provides functions to parse various structured data formats into an internal Go representation that can later be reformatted by the formatter package.
package parser

import (
	"fmt"
	"sort"
)

type parserFunc func([]byte) (interface{}, error)

type parser struct {
	parse   parserFunc
	prefers []string
}

var parsers = make(map[string]parser)

// Names returns a sorted list of valid options for the "format" parameter of Parse
func Names() []string {
	names := make([]string, 0, len(parsers))

	for name, _ := range parsers {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Parse takes serialised data and parses it into a Go interface{}.
// The format can be supplied explicitly or can be set to "auto" in which case Parse will attempt to automatically identify the serialisation format.
func Parse(input []byte, format string) (interface{}, error) {
	parser, ok := parsers[format]

	if !ok {
		return nil, fmt.Errorf("No such parser: %s", format)
	}

	return parser.parse(input)
}
