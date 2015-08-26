package formatters

import (
	"fmt"

	"github.com/nytlabs/mxj"
	"github.com/stilvoid/please/util"
)

func JSON(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

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