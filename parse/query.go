package parse

import "net/url"

func Query(input []byte) (interface{}, error) {
	result, err := url.ParseQuery(string(input))

	return map[string][]string(result), err
}
