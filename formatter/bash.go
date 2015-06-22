package formatter

import (
	"fmt"
	"reflect"
	"strings"
)

func wrapObj(in interface{}) string {
	out := Bash(in)
	out = strings.Replace(out, "\\", "\\\\", -1)
	out = strings.Replace(out, "\"", "\\\"", -1)
	out = strings.Replace(out, "\n", "\\n", -1)
	out = strings.Replace(out, "$", "\\$", -1)
	out = fmt.Sprintf("\"%s\"", out)

	return out
}

func Bash(in interface{}) (out string) {

	if in == nil {
		return ""
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		vv := in.(map[string]interface{})

		parts := make([]string, len(vv))

		i := 0

		for key, value := range vv {
			parts[i] = fmt.Sprintf("[%s]=%s", key, wrapObj(value))
			i++
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " "))
	case reflect.Array, reflect.Slice:
		parts := make([]string, val.Len())

		i := 0

		for index := 0; index < val.Len(); index++ {
			value := val.Index(index).Interface()

			parts[i] = fmt.Sprintf("[%d]=%s", index, wrapObj(value))
			i++
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " "))
	default:
		return fmt.Sprint(in)
	}
}
