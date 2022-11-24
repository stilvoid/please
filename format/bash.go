package format

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/stilvoid/please/internal"
)

func wrapObj(in string) string {
	out := strings.Replace(in, `\`, `\\`, -1)
	out = strings.Replace(out, `"`, `\"`, -1)
	out = strings.Replace(out, "`", "\\`", -1)
	out = strings.Replace(out, "\n", `\n`, -1)
	out = strings.Replace(out, `$`, `\$`, -1)

	return "\"" + out + "\""
}

func formatBashInternal(in interface{}, buf *bytes.Buffer) {
	if in == nil {
		return
	}

	switch v := in.(type) {
	case map[string]interface{}:
		keys := internal.SortedKeys(v)

		buf.WriteByte('(')

		for i, key := range keys {
			var innerBuf bytes.Buffer

			formatBashInternal(v[key.(string)], &innerBuf)

			buf.WriteByte('[')
			buf.WriteString(key.(string))
			buf.WriteString("]=")
			buf.WriteString(wrapObj(innerBuf.String()))

			if i != len(keys)-1 {
				buf.WriteByte(' ')
			}
		}

		buf.WriteByte(')')
	default:
		buf.WriteString(fmt.Sprint(in))
	}
}

func formatBash(in interface{}) (string, error) {
	in = internal.ArraysToMaps(in)
	in = internal.ForceStringKeys(in)

	var buf bytes.Buffer

	formatBashInternal(in, &buf)

	return buf.String(), nil
}
