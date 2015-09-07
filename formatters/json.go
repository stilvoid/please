package formatters

import (
	"encoding/json"

	"github.com/stilvoid/please/common"
)

func formatJSON(in interface{}) (string, error) {
	in = common.ForceStringKeys(in)

	bytes, err := json.MarshalIndent(in, "", "  ")

	return string(bytes), err
}
