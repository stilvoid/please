package parser

import (
	"encoding/json"
)

func parseJson(input []byte) (interface{}, error) {
	var parsed interface{}

	err := json.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	parsers["json"] = parser{
		parse: parseJson,
	}
}
