package parser

import (
	"github.com/clbanning/x2j"
)

func Xml(input []byte) (interface{}, error) {
	parsed := make(map[string]interface{})

	err := x2j.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	Parsers["xml"] = parser{
		parse:   Xml,
		prefers: []string{"json"},
	}
}
