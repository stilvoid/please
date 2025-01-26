package format

import (
	"bytes"
	"fmt"
	"maps"
	"slices"
	"sort"
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

func formatBashInternal(in any, buf *bytes.Buffer) {
	if in == nil {
		return
	}

	switch v := in.(type) {
	case map[string]any:
		keys := slices.Collect(maps.Keys(v))
		sort.Strings(keys)

		buf.WriteByte('(')

		for i, key := range keys {
			var innerBuf bytes.Buffer

			formatBashInternal(v[key], &innerBuf)

			buf.WriteByte('[')
			buf.WriteString(key)
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

func Bash(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{
		MapArrays:  true,
		StringKeys: true,
	})

	var buf bytes.Buffer

	formatBashInternal(in, &buf)

	return buf.String(), nil
}
