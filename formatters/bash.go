package formatters

import (
	"fmt"
	"strings"

	"github.com/stilvoid/please/util"
)

func wrapObj(in interface{}) string {
	out := formatBashInternal(in)
	out = strings.Replace(out, "\\", "\\\\", -1)
	out = strings.Replace(out, "\"", "\\\"", -1)
	out = strings.Replace(out, "`", "\\`", -1)
	out = strings.Replace(out, "\n", "\\n", -1)
	out = strings.Replace(out, "$", "\\$", -1)
	out = fmt.Sprintf("\"%s\"", out)

	return out
}

func formatBashInternal(in interface{}) string {
	if in == nil {
		return ""
	}

	switch v := in.(type) {
	case map[string]interface{}:
		keys := util.SortedKeys(v)

		parts := make([]string, len(v))

		for i, key := range keys {
			parts[i] = fmt.Sprintf("[%s]=%s", key, wrapObj(v[key.(string)]))
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " "))
	default:
		return fmt.Sprint(in)
	}
}

func formatBash(in interface{}) (string, error) {
	in = util.ArraysToMaps(in)
	in = util.ForceStringKeys(in)

	return formatBashInternal(in), nil
}
