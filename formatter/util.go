package formatter

import (
	"fmt"
	"reflect"
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
