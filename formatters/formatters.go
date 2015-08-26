package formatters

import (
	"fmt"
	"sort"
)

// Type Formatter is a function that takes an interface{} and attempts to format it as a string
type Formatter func(interface{}) (string, error)

var formatters = make(map[string]Formatter)

func init() {
	Register("bash", formatBash)
	Register("dot", formatDot)
	Register("json", formatJSON)
	Register("xml", formatXML)
	Register("yaml", formatYAML)
}

// Names returns a sorted list of valid options for the "format" parameter of Format
func Names() []string {
	names := make([]string, 0, len(formatters))

	for name, _ := range formatters {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Registerassigns a Formatter function to a name. If the name has already been registered, an error will be returned.
func Register(name string, formatter Formatter) error {
	if _, ok := formatters[name]; ok {
		return fmt.Errorf("Formatter '%s' already exists", name)
	}

	formatters[name] = formatter

	return nil
}

// Getreturns a Formatter function by name. If the named formatter is not found, an error will be returned.
func Get(name string) (Formatter, error) {
	formatter, ok := formatters[name]

	if !ok {
		return nil, fmt.Errorf("no such formatter: %s", name)
	}

	return formatter, nil
}
