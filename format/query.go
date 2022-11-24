package format

import (
	"fmt"
	"net/url"

	"github.com/stilvoid/please/internal"
)

func formatQueryInternal(in interface{}) string {
	if in == nil {
		return ""
	}

	inMap, ok := in.(map[string]interface{})

	if !ok {
		return fmt.Sprint(in)
	}

	var output url.Values = make(map[string][]string)

	for key, value := range inMap {
		switch value.(type) {
		case map[string]interface{}:
			result := formatQueryInternal(value)

			output.Add(key, result)
		case nil:
			output.Add(key, "")
		default:
			output.Add(key, fmt.Sprint(value))
		}
	}

	return output.Encode()
}

func formatQuery(in interface{}) (string, error) {
	in = internal.ArraysToMaps(in)
	in = internal.ForceStringKeys(in)

	return formatQueryInternal(in), nil
}
