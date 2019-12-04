package formatters

import (
	"encoding/json"
)

func formatJSON(in interface{}) (string, error) {
	in = forceStringKeys(in)

	bytes, err := json.MarshalIndent(in, "", "  ")

	return string(bytes), err
}
