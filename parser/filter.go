package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Filter takes structured data and returns just the portion of it that matches the provided path using dot notation.
// The path may contain asterisks, in which case all entries will match at that level of the path
// e.g. products.*.id will return all product ids
func Filter(in interface{}, path string) (interface{}, error) {
	if path == "" {
		return in, nil
	}

	splitPath := strings.SplitN(path, ".", 2)

	thisPath := splitPath[0]
	var nextPath string

	if len(splitPath) > 1 {
		nextPath = splitPath[1]
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		if thisPath == "*" {
			var err error
			out := make([]interface{}, val.Len())

			for i, key := range val.MapKeys() {
				out[i], err = Filter(val.MapIndex(key).Interface(), nextPath)

				if err != nil {
					return nil, err
				}
			}

			return out, nil
		}

		for _, key := range val.MapKeys() {
			if fmt.Sprint(key.Interface()) == thisPath {
				value := val.MapIndex(key).Interface()
				return Filter(value, nextPath)
			}
		}

		break
	case reflect.Array, reflect.Slice:
		if thisPath == "*" {
			var err error
			out := make([]interface{}, val.Len())

			for i := 0; i < val.Len(); i++ {
				out[i], err = Filter(val.Index(i).Interface(), nextPath)

				if err != nil {
					return nil, err
				}
			}

			return out, nil
		}

		index, err := strconv.Atoi(thisPath)

		if err != nil || index < 0 || index >= val.Len() {
			break
		}

		return Filter(val.Index(index).Interface(), nextPath)
	}

	return nil, fmt.Errorf("key does not exist: %s", thisPath)
}
