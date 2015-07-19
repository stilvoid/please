package parser

import (
	"fmt"
)

type parserFunc func([]byte) (interface{}, error)

type parser struct {
	parse   parserFunc
	prefers []string
}

var Parsers = make(map[string]parser)

func Parse(input []byte, format string) (interface{}, error) {
	parser, ok := Parsers[format]

	if !ok {
		return nil, fmt.Errorf("No such parser: %s", format)
	}

	return parser.parse(input)
}
