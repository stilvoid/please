package parsers

import "github.com/clbanning/x2j"

func XML(input []byte) (interface{}, error) {
	parsed := make(map[string]interface{})

	err := x2j.Unmarshal(input, &parsed)

	return parsed, err
}
