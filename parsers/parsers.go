package parsers

import (
	"fmt"
	"sort"
)

// Type Parser is a function that takes a byte slice and attempts to parse it into a structure format in an interface{}
type Parser func([]byte) (interface{}, error)

var parsers = make(map[string]Parser)

func init() {
	Register("csv", CSV)
	Register("html", HTML)
	Register("json", JSON)
	Register("mime", MIME)
	Register("xml", XML)
	Register("yaml", YAML)
}

// Names returns a sorted list of valid options for the "format" parameter of Parse
func Names() []string {
	names := make([]string, 0, len(parsers))

	for name, _ := range parsers {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Registerassigns a Parser function to a name. If the name has already been registered, an error will be returned.
func Register(name string, parser Parser) error {
	if _, ok := parsers[name]; ok {
		return fmt.Errorf("Parser %s already exists", name)
	}

	parsers[name] = parser

	return nil
}

// Getreturns a Parser function by name. If the named parser is not found, an error will be returned.
func Get(name string) (Parser, error) {
	parser, ok := parsers[name]

	if !ok {
		return nil, fmt.Errorf("No such parser: %s", name)
	}

	return parser, nil
}
