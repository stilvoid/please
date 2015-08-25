// Package formatter provides function to format structured data into a variety of serialisation formats
package formatter

import (
	"fmt"
	"sort"
)

// Type Formatter is a function that takes an interface{} and attempts to format it as a string
type Formatter func(interface{}) (string, error)

var formatters = make(map[string]Formatter)

// Names returns a sorted list of valid options for the "format" parameter of Format
func Names() []string {
	names := make([]string, 0, len(formatters))

	for name, _ := range formatters {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Get returns a Formatter function by name. If the named formatter is not found, an error will be returned.
func Get(name string) (Formatter, error) {
	formatter, ok := formatters[name]

	if !ok {
		return nil, fmt.Errorf("no such formatter: %s", name)
	}

	return formatter, nil
}
