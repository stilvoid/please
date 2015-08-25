// Package parser provides functions to parse various structured data formats into an internal Go representation that can later be reformatted by the formatter package.
package parser

import (
	"fmt"
	"sort"
)

// Type Parser is a function that takes a byte slice and attempts to parse it into a structure format in an interface{}
type Parser func([]byte) (interface{}, error)

type parser struct {
	parse   Parser
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

// Get returns a Parser function by name. If the named parser is not found, an error will be returned.
func Get(name string) (Parser, error) {
	parser, ok := parsers[name]

	if !ok {
		return nil, fmt.Errorf("No such parser: %s", name)
	}

	return parser.parse, nil
}
