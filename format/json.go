package format

import (
	"encoding/json"

	"github.com/stilvoid/please/internal"
)

func Json(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{StringKeys: true})

	bytes, err := json.MarshalIndent(in, "", "  ")

	return string(bytes), err
}
