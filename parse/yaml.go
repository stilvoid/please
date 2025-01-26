package parse

import "gopkg.in/yaml.v2"

func Yaml(input []byte) (any, error) {
	var parsed any

	err := yaml.Unmarshal(input, &parsed)

	return parsed, err
}
