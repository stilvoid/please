package parsers

import "encoding/json"

func parseJSON(input []byte) (interface{}, error) {
	var parsed interface{}

	err := json.Unmarshal(input, &parsed)

	return parsed, err
}
