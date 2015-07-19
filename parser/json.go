package parser

import (
	"encoding/json"
)

func Json(input []byte) (interface{}, error) {
	var parsed interface{}

	err := json.Unmarshal(input, &parsed)

	return parsed, err
}

func init() {
	Parsers["json"] = parser{
		parse: Json,
	}
}
