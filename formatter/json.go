package formatter

import (
	"fmt"

	"github.com/nytlabs/mxj"
)

func formatJSON(in interface{}) (string, error) {
	inMap, ok := in.(map[string]interface{})

	if !ok {
		return fmt.Sprintf("\"%s\"", in), nil
	}

	m := mxj.Map(inMap)

	bytes, err := m.JsonIndent("", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func init() {
	formatters["json"] = formatJSON
}
