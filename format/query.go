package format

import (
	"fmt"
	"net/url"

	"github.com/stilvoid/please/internal"
)

func formatQueryInternal(in any) string {
	if in == nil {
		return ""
	}

	inMap, ok := in.(map[string]any)

	if !ok {
		return fmt.Sprint(in)
	}

	var output url.Values = make(map[string][]string)

	for key, value := range inMap {
		switch value.(type) {
		case map[string]any:
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

func Query(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{
		MapArrays:  true,
		StringKeys: true,
	})

	return formatQueryInternal(in), nil
}
