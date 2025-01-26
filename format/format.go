package format

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

var formatters = map[string]func(any) (string, error){
	"bash":  Bash,
	"dot":   Dot,
	"json":  Json,
	"xml":   Xml,
	"yaml":  Yaml,
	"query": Query,
	"toml":  Toml,
}

// Formatters are the names of data formats that please can output
var Formatters []string

func init() {
	Formatters = slices.Collect(maps.Keys(formatters))
	sort.Strings(Formatters)
}

// Format converts input data into format as a byte array
func Format(format string, input any) (string, error) {
	formatter, ok := formatters[format]
	if !ok {
		return "", fmt.Errorf("Invalid formatter: %s", format)
	}

	return formatter(input)
}
