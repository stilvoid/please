package util

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func ForceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		new_map := make(map[string]interface{}, val.Len())

		for _, key := range val.MapKeys() {
			string_key := fmt.Sprint(key.Interface())
			new_map[string_key] = ForceStringKeys(val.MapIndex(key).Interface())
		}

		return new_map
	case reflect.Array, reflect.Slice:
		new_slice := make([]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			new_slice[i] = ForceStringKeys(value)
		}

		return new_slice
	default:
		return in
	}
}

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

func SortKeys(in interface{}) []string {
	v := reflect.ValueOf(in)

	if v.Kind() != reflect.Map {
		panic("Not a map!")
	}

	vkeys := v.MapKeys()

	keys := make([]string, 0, len(vkeys))

	for _, key := range vkeys {
		keys = append(keys, fmt.Sprint(key.Interface()))
	}

	sort.Strings(keys)

	return keys
}
