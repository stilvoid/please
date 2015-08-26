package parsers

import "gopkg.in/yaml.v2"

func YAML(input []byte) (interface{}, error) {
	var parsed interface{}

	err := yaml.Unmarshal(input, &parsed)

	return parsed, err
}
