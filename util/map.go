package util

import (
	"fmt"
	"reflect"
)

// ForceStringKeys creates a copy of the provided interface{}, with all maps changed to have string keys for use by serialisers that expect string keys
// This is particularly useful for formatters where the target serialisation format only allows string keys
func ForceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		newMap := make(map[string]interface{}, val.Len())

		var stringKey string

		for _, key := range val.MapKeys() {
			if reflect.TypeOf(key.Interface()) == nil {
				stringKey = "null"
			} else {
				stringKey = fmt.Sprint(key.Interface())
			}

			newMap[stringKey] = ForceStringKeys(val.MapIndex(key).Interface())
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
