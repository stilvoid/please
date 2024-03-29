package format

import (
	"encoding/json"

	"github.com/stilvoid/please/internal"
)

func Json(in interface{}) (string, error) {
	in = internal.ForceStringKeys(in)

	bytes, err := json.MarshalIndent(in, "", "  ")

	return string(bytes), err
}
