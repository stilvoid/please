package parse

import "net/url"

func parseQuery(input []byte) (interface{}, error) {
	result, err := url.ParseQuery(string(input))

	return map[string][]string(result), err
}
