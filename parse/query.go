package parse

import "net/url"

func Query(input []byte) (any, error) {
	result, err := url.ParseQuery(string(input))

	return map[string][]string(result), err
}
