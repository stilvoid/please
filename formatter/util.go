package formatter

import (
	"fmt"
	"reflect"
)

func forceStringKeys(in interface{}) interface{} {
	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		newMap := make(map[string]interface{}, val.Len())

		for _, key := range val.MapKeys() {
			stringKey := fmt.Sprint(key.Interface())
			newMap[stringKey] = forceStringKeys(val.MapIndex(key).Interface())
		}

		return newMap
	case reflect.Array, reflect.Slice:
		newSlice := make([]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			newSlice[i] = forceStringKeys(value)
		}

		return newSlice
	default:
		return in
	}
}
