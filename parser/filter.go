package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Filter takes structured data and returns just the portion of it that matches the provided path using dot notation.
// e.g. data.header.1
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
		for _, key := range val.MapKeys() {
			if fmt.Sprint(key.Interface()) == thisPath {
				value := val.MapIndex(key).Interface()
				return Filter(value, nextPath)
			}
		}

		break
	case reflect.Array, reflect.Slice:
		index, err := strconv.Atoi(thisPath)

		if err != nil || index < 0 || index >= val.Len() {
			break
		}

		return Filter(val.Index(index).Interface(), nextPath)
	}

	return nil, fmt.Errorf("Key does not exist: %s", thisPath)
}
