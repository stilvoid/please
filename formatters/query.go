package formatters

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/stilvoid/please/util"
)

// TODO: Tidy this up; it's horrid.
func formatQuery(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	inMap, ok := in.(map[string]interface{})

	if !ok {
		return "", fmt.Errorf("query formatter expects a map as input")
	}

	var output url.Values = make(map[string][]string)

	for key, value := range inMap {
		val := reflect.ValueOf(value)

		switch val.Kind() {
		case reflect.Map:
			return "", fmt.Errorf("query formatter cannot deal with nested values")
		case reflect.Array, reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				iVal := val.Index(i)

				switch iVal.Kind() {
				case reflect.Map, reflect.Array, reflect.Slice:
					return "", fmt.Errorf("query formatter cannot deal with nested values")
				default:
					output.Add(key, fmt.Sprint(iVal.Interface()))
				}
			}
		default:
			output.Add(key, fmt.Sprint(value))
		}
	}

	return output.Encode(), nil
}
