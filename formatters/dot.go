package formatters

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/stilvoid/please/common"
)

type node struct {
	name  string
	label string
}

type link struct {
	left  string
	right string
}

func wrap(in string) string {
	out := strings.Replace(in, "\\", "\\\\", -1)
	out = strings.Replace(out, "\"", "\\\"", -1)
	out = strings.Replace(out, "\n", "\\n", -1)
	out = fmt.Sprintf("\"%s\"", out)

	return out
}

func flatten(in interface{}, currentPath string) ([]node, []link) {
	var nodes []node
	var links []link

	if in == nil {
		return nodes, links
	}

	val := reflect.ValueOf(in)

	switch val.Kind() {
	case reflect.Map:
		vv := in.(map[string]interface{})

		for key, value := range vv {
			target := currentPath + "-" + key

			nodes = append(nodes, node{
				name:  target,
				label: key,
			})

			if currentPath != "" {
				links = append(links, link{
					left:  currentPath,
					right: target,
				})
			}

			newNodes, newLinks := flatten(value, target)

			nodes = append(nodes, newNodes...)
			links = append(links, newLinks...)
		}

		return nodes, links
	case reflect.Array, reflect.Slice:
		for index := 0; index < val.Len(); index++ {
			value := val.Index(index).Interface()

			target := currentPath + "-" + fmt.Sprint(index)

			nodes = append(nodes, node{
				name:  target,
				label: fmt.Sprintf("[%d]", index),
			})

			if currentPath != "" {
				links = append(links, link{
					left:  currentPath,
					right: target,
				})
			}

			newNodes, newLinks := flatten(value, target)

			nodes = append(nodes, newNodes...)
			links = append(links, newLinks...)
		}

		return nodes, links
	default:
		target := currentPath + "=content"

		nodes = append(nodes, node{
			name:  target,
			label: fmt.Sprint(in),
		})

		links = append(links, link{
			left:  currentPath,
			right: target,
		})

		return nodes, links
	}
}

func formatDot(in interface{}) (string, error) {
	in = common.ForceStringKeys(in)

	nodes, links := flatten(in, "root")

	nodes = append(nodes, node{
		name:  "root",
		label: "[Root]",
	})

	var buf bytes.Buffer

	for _, node := range nodes {
		buf.WriteString(fmt.Sprintf("%s [label=%s];\n", wrap(node.name), wrap(node.label)))
	}

	for _, link := range links {
		buf.WriteString(fmt.Sprintf("%s -- %s;\n", wrap(link.left), wrap(link.right)))
	}

	return fmt.Sprintf("graph{\n%s}", buf.String()), nil
}
