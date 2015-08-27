package formatters

import (
	"encoding/json"

	"github.com/stilvoid/please/util"
)

func formatJSON(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	bytes, err := json.MarshalIndent(in, "", "  ")

	return string(bytes), err
}
