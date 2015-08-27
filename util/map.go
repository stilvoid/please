package util

import (
	"fmt"
	"reflect"
	"sort"
)

// toString wraps fmt.Sprint except that it converts nil to "null"
func toString(in interface{}) string {
	if in == nil {
		return "null"
	}

	return fmt.Sprint(in)
}

// ForceStringKeys creates a copy of the provided interface{}, with all maps changed to have string keys for use by serialisers that expect string keys
// This is useful for formatters where the target serialisation format only allows string keys
func ForceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		newMap := make(map[string]interface{}, val.Len())

		for _, keyVal := range val.MapKeys() {
			key := toString(keyVal.Interface())
			value := val.MapIndex(keyVal).Interface()

			newMap[key] = ForceStringKeys(value)
		}

		return newMap
	case reflect.Array, reflect.Slice:
		newSlice := make([]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			newSlice[i] = ForceStringKeys(value)
		}

		return newSlice
	default:
		return in
	}
}

// ArraysToMaps creates a copy of the provided interface{}, with all arrays converted into maps where the keys are the array indices, starting at 0.
// This is useful for formatters where the target serialisation format does not have a means of representing arrays
func ArraysToMaps(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		newMap := make(map[interface{}]interface{}, val.Len())

		for _, key := range val.MapKeys() {
			value := val.MapIndex(key).Interface()

			newMap[key.Interface()] = ArraysToMaps(value)
		}

		return newMap
	case reflect.Array, reflect.Slice:
		newMap := make(map[interface{}]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()

			newMap[interface{}(i)] = ArraysToMaps(value)
		}

		return newMap
	default:
		return in
	}
}

func SortedKeys(in interface{}) []interface{} {
	val := reflect.ValueOf(in)

	if val.Kind() != reflect.Map {
		panic("SortedKeys only works on maps")
	}

	stringKeys := make([]string, val.Len())
	stringKeysMap := make(map[string]interface{}, val.Len())

	for i, key := range val.MapKeys() {
		stringKey := toString(key.Interface())

		stringKeys[i] = stringKey

		stringKeysMap[stringKey] = key.Interface()
	}

	sort.Strings(stringKeys)

	outKeys := make([]interface{}, val.Len())

	for i, key := range stringKeys {
		outKeys[i] = stringKeysMap[key]
	}

	return outKeys
}
