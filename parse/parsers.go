package parse

import (
	"fmt"
	"sort"
)

// Type Parser is a function that takes a byte slice and attempts to parse it into a structure format in an any
type Parser func([]byte) (any, error)

var parsers = make(map[string]Parser)

func init() {
	Register("csv", Csv)
	Register("html", Html)
	Register("json", Json)
	Register("mime", Mime)
	Register("xml", Xml)
	Register("yaml", Yaml)
	Register("query", Query)
	Register("toml", Toml)
}

// Names returns a sorted list of valid options for the "format" parameter of Parse
func Names() []string {
	names := make([]string, 0, len(parsers))

	for name := range parsers {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Register assigns a Parser function to a name. If the name has already been registered, an error will be returned.
func Register(name string, parser Parser) error {
	if _, ok := parsers[name]; ok {
		return fmt.Errorf("parser %s already exists", name)
	}

	parsers[name] = parser

	return nil
}

// Get returns a Parser function by name. If the named parser is not found, an error will be returned.
func Get(name string) (Parser, error) {
	parser, ok := parsers[name]

	if !ok {
		return nil, fmt.Errorf("no such parser: %s", name)
	}

	return parser, nil
}
