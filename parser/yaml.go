package parser

import (
	"gopkg.in/yaml.v2"
)

func Yaml(input []byte) (interface{}, error) {
	var parsed interface{}

	err := yaml.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	Parsers["yaml"] = parser{
		parse:   Yaml,
		prefers: []string{"json", "xml"},
	}
}
