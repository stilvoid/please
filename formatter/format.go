// Package formatter provides function to format structured data into a variety of serialisation formats
package formatter

import (
	"fmt"
	"sort"

	"github.com/stilvoid/please/util"
)

type formatterFunc func(interface{}) (string, error)

var formatters = make(map[string]formatterFunc)

// Names returns a sorted list of valid options for the "format" parameter of Format
func Names() []string {
	names := make([]string, 0, len(formatters))

	for name, _ := range formatters {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Parse takes a Go interface{} and serialises it into the format specified.
func Format(input interface{}, format string) (string, error) {
	formatter, ok := formatters[format]

	if !ok {
		return "", fmt.Errorf("no such formatter: %s", format)
	}

	if format != "yaml" {
		// Pretty much everything hates non-string keys :S
		input = util.ForceStringKeys(input)
	}

	return formatter(input)
}
