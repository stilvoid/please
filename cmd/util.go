package cmd

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func forceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		new_map := make(map[string]interface{}, val.Len())

		for _, key := range val.MapKeys() {
			string_key := fmt.Sprint(key.Interface())
			new_map[string_key] = forceStringKeys(val.MapIndex(key).Interface())
		}

		return new_map
	case reflect.Array, reflect.Slice:
		new_slice := make([]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			new_slice[i] = forceStringKeys(value)
		}

		return new_slice
	default:
		return in
	}
}

func filter(in interface{}, path string) interface{} {
	if path == "" {
		return in
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
		vv := in.(map[string]interface{})

		next, ok := vv[this_path]

		if !ok {
			break
		}

		return filter(next, next_path)
	case reflect.Array, reflect.Slice:
		index, err := strconv.Atoi(this_path)

		if err != nil || index < 0 || index >= val.Len() {
			break
		}

		return filter(val.Index(index).Interface(), next_path)
	}

	fmt.Fprintf(os.Stderr, "Key does not exist %s\n", this_path)
	os.Exit(1)

	return nil
}

func sortKeys(in interface{}) []string {
	v := reflect.ValueOf(in)

	if v.Kind() != reflect.Map {
		panic("Not a map!")
	}

	vkeys := v.MapKeys()

	keys := make([]string, 0, len(vkeys))

	for _, key := range vkeys {
		keys = append(keys, key.String())
	}

	sort.Strings(keys)

	return keys
}
