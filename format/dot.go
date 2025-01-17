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

func wrap(in string) string {
	out := strings.Replace(in, `\`, `\\`, -1)
	out = strings.Replace(out, `"`, `\"`, -1)
	out = strings.Replace(out, "\n", `\n`, -1)

	return "\"" + out + "\""
}

func dotNode(name string, label string) string {
	return wrap(name) + " [label=" + wrap(label) + "];\n"
}

func dotLink(left string, right string, note ...string) string {
	if len(note) > 0 {
		return wrap(left) + " -- " + wrap(right) + " [label=" + wrap(fmt.Sprint(note)) + "];\n"
	}

	return wrap(left) + " -- " + wrap(right) + ";\n"
}

func flatten(in any, parent string, name string, buf *bytes.Buffer) {
	switch vv := in.(type) {
	case map[string]any:
		if parent != "" {
			buf.WriteString(dotLink(parent, name))
		}

		buf.WriteString(dotNode(name, "[map]"))

		parent = name

		i := 0

		keys := slices.Collect(maps.Keys(vv))
		sort.Strings(keys)

		for _, key := range keys {
			value := vv[key]

			target := parent + "-map-" + fmt.Sprint(i)

			buf.WriteString(dotLink(parent, target))
			buf.WriteString(dotNode(target, key))

			contentTarget := target + "=content"

			flatten(value, target, contentTarget, buf)

			i++
		}
	case []any:
		if parent != "" {
			buf.WriteString(dotLink(parent, name))
		}

		buf.WriteString(dotNode(name, "[array]"))

		parent = name

		for i, value := range vv {
			target := name + "-array-" + fmt.Sprint(i)

			flatten(value, parent, target, buf)
		}
	default:
		if parent != "" {
			buf.WriteString(dotLink(parent, name))
		}

		buf.WriteString(dotNode(name, fmt.Sprint(in)))
	}
}

func Dot(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{StringKeys: true})

	var buf bytes.Buffer

	buf.WriteString("graph{\n")

	flatten(in, "", "root", &buf)

	buf.WriteByte('}')

	return buf.String(), nil
}
