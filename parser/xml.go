package parser

import (
	"github.com/clbanning/x2j"
)

func parseXml(input []byte) (interface{}, error) {
	parsed := make(map[string]interface{})

	err := x2j.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	parsers["xml"] = parser{
		parse:   parseXml,
		prefers: []string{"json"},
	}
}
