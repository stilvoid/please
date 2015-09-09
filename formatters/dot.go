package formatters

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/stilvoid/please/common"
)

var escapeRe *regexp.Regexp
var newlineRe *regexp.Regexp

func init() {
	escapeRe = regexp.MustCompile(`(\\|\")`)
	newlineRe = regexp.MustCompile(`\n`)
}

func wrap(in string) string {
	out := strings.Replace(in, `\`, `\\`, -1)
	out = strings.Replace(out, `"`, `\"`, -1)
	out = strings.Replace(out, "\n", `\n`, -1)

	return "\"" + out + "\""
}

func dotNode(name string, label string) string {
	return wrap(name) + " [label=" + wrap(label) + "];\n"
}

func dotLink(left string, right string) string {
	return wrap(left) + " -- " + wrap(right) + ";\n"
}

func flatten(in interface{}, currentPath string, buf *bytes.Buffer) {
	if in == nil {
		return
	}

	switch vv := in.(type) {
	case map[string]interface{}:
		for key, value := range vv {
			target := currentPath + "-" + key

			buf.WriteString(dotNode(target, key))

			if currentPath != "" {
				buf.WriteString(dotLink(currentPath, target))
			}

			flatten(value, target, buf)
		}
	default:
		target := currentPath + "=content"

		buf.WriteString(dotNode(target, fmt.Sprint(in)))

		buf.WriteString(dotLink(currentPath, target))
	}
}

func formatDot(in interface{}) (string, error) {
	in = common.ArraysToMaps(in)
	in = common.ForceStringKeys(in)

	var buf bytes.Buffer

	buf.WriteString("graph{\n")

	buf.WriteString(dotNode("root", "[Root]"))

	flatten(in, "root", &buf)

	buf.WriteByte('}')

	return buf.String(), nil
}
