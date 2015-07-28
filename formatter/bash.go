package formatter

import (
	"fmt"
	"reflect"
	"strings"
)

func wrapObj(in interface{}) string {
	out := formatBash(in)
	out = strings.Replace(out, "\\", "\\\\", -1)
	out = strings.Replace(out, "\"", "\\\"", -1)
	out = strings.Replace(out, "\n", "\\n", -1)
	out = strings.Replace(out, "$", "\\$", -1)
	out = fmt.Sprintf("\"%s\"", out)

	return out
}

func formatBash(in interface{}) (out string) {

	if in == nil {
		return ""
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		parts := make([]string, val.Len())

		for i, key := range val.MapKeys() {
			value := val.MapIndex(key).Interface()
			parts[i] = fmt.Sprintf("[%s]=%s", key.String(), wrapObj(value))
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " "))
	case reflect.Array, reflect.Slice:
		parts := make([]string, val.Len())

		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()

			parts[i] = fmt.Sprintf("[%d]=%s", i, wrapObj(value))
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " "))
	default:
		return fmt.Sprint(in)
	}
}

func init() {
	formatters["bash"] = formatBash
}
