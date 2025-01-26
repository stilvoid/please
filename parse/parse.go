package parse

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

var parsers = map[string]func([]byte) (any, error){
	"csv":   Csv,
	"html":  Html,
	"json":  Json,
	"mime":  Mime,
	"xml":   Xml,
	"yaml":  Yaml,
	"query": Query,
	"toml":  Toml,
}

// Parsers are the names of data formats that please can read
var Parsers []string

func init() {
	Parsers = slices.Collect(maps.Keys(parsers))
	sort.Strings(Parsers)
}

// Parse reads input as format and returns data
func Parse(format string, input []byte) (any, error) {
	parser, ok := parsers[format]
	if !ok {
		return nil, fmt.Errorf("Invalid parser: %s", format)
	}

	return parser(input)
}
