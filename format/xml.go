package format

import (
	"bytes"
	"fmt"
	"maps"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/stilvoid/please/internal"
)

const INDENT = "  "

const badTagChars = "(^[-.0-9]|[!\"#$%&'()*+,/;<=>?@\\[\\\\]^`{|}~]|\\s)"

var badTagRe *regexp.Regexp

func init() {
	badTagRe = regexp.MustCompile(badTagChars)
}

func xmlTag(in string, attributes map[string]string) string {
	out := "<" + badTagRe.ReplaceAllLiteralString(in, "_")

	for key, value := range attributes {
		out += " " + key + "=\"" + strings.Replace(value, "\"", "\\\"", -1) + "\""
	}

	out += ">"

	return out
}

func xmlCloseTag(in string) string {
	return "</" + badTagRe.ReplaceAllLiteralString(in, "_") + ">"
}

func xmlWrap(in any) string {
	out := fmt.Sprint(in)
	out = strings.Replace(out, "&", "&amp;", -1)
	out = strings.Replace(out, "<", "&lt;", -1)
	out = strings.Replace(out, ">", "&gt;", -1)

	return out
}

func doIndent(indent int, buf *bytes.Buffer) {
	for i := 0; i < indent; i++ {
		buf.WriteString(INDENT)
	}
}

func formatXMLInternal(in any, parent string, indent int, buf *bytes.Buffer) {
	switch v := in.(type) {
	case map[string]any:
		attributes := make(map[string]string)
		children := make(map[string]any)
		text := make([]any, 0)

		keys := slices.Collect(maps.Keys(v))
		sort.Strings(keys)

		// Gather attributes
		for _, key := range keys {
			value := v[key]

			if key[0] == '-' {
				attributes[key[1:]] = fmt.Sprint(value)
			} else if key == "#text" {
				val := reflect.ValueOf(value)

				switch val.Kind() {
				case reflect.Slice, reflect.Array:
					for i := 0; i < val.Len(); i++ {
						text = append(text, val.Index(i).Interface())
					}
				default:
					text = append(text, value)
				}
			} else {
				children[key] = value
			}
		}

		// Write the tag
		doIndent(indent, buf)
		buf.WriteString(xmlTag(parent, attributes))

		// Write text and children
		if len(text) == 1 && len(children) == 0 {
			buf.WriteString(xmlWrap(text[0]))
		} else if len(text) > 1 || len(children) > 0 {
			buf.WriteString("\n")

			for _, line := range text {
				doIndent(indent+1, buf)
				buf.WriteString(xmlWrap(line))
				buf.WriteString("\n")
			}

			keys = slices.Collect(maps.Keys(children))
			sort.Strings(keys)

			for _, key := range keys {
				value := children[key]

				if _, ok := value.([]any); ok && key == "root" {
					key = "tag"
				}

				formatXMLInternal(value, key, indent+1, buf)
			}

			doIndent(indent, buf)
		}

		// Close up
		buf.WriteString(xmlCloseTag(parent))
		buf.WriteString("\n")
	case []any:
		for _, value := range v {
			if _, ok := value.([]any); ok {
				doIndent(indent, buf)
				buf.WriteString(xmlTag(parent, nil))
				buf.WriteString("\n")
				formatXMLInternal(value, "tag", indent+1, buf)
				doIndent(indent, buf)
				buf.WriteString(xmlCloseTag(parent))
				buf.WriteString("\n")
			} else {
				formatXMLInternal(value, parent, indent, buf)
			}
		}
	default:
		doIndent(indent, buf)
		buf.WriteString(xmlTag(parent, nil))
		buf.WriteString(xmlWrap(v))
		buf.WriteString(xmlCloseTag(parent))
		buf.WriteString("\n")
	}
}

func Xml(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{StringKeys: true})

	if _, ok := in.([]any); ok {
		in = map[string]any{
			"root": in,
		}
	}

	var buf bytes.Buffer

	formatXMLInternal(in, "root", 0, &buf)

	return strings.TrimSpace(buf.String()), nil
}
