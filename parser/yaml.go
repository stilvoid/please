package parser

import (
	"gopkg.in/yaml.v2"
)

func parseYaml(input []byte) (interface{}, error) {
	var parsed interface{}

	err := yaml.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	parsers["yaml"] = parser{
		parse:   parseYaml,
		prefers: []string{"json", "xml"},
	}
}
