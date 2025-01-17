package parse

import "encoding/json"

func Json(input []byte) (any, error) {
	var parsed any

	err := json.Unmarshal(input, &parsed)

	return parsed, err
}
