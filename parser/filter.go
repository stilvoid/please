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

	split_path := strings.SplitN(path, ".", 2)

	this_path := split_path[0]
	var next_path string

	if len(split_path) > 1 {
		next_path = split_path[1]
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			if fmt.Sprint(key.Interface()) == this_path {
				value := val.MapIndex(key).Interface()
				return Filter(value, next_path)
			}
		}

		break
	case reflect.Array, reflect.Slice:
		index, err := strconv.Atoi(this_path)

		if err != nil || index < 0 || index >= val.Len() {
			break
		}

		return Filter(val.Index(index).Interface(), next_path)
	}

	return nil, fmt.Errorf("Key does not exist: %s", this_path)
}
