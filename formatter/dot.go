package formatter

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
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

func Dot(in interface{}) (out string) {
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

	return fmt.Sprintf("graph{\n%s}", buf.String())
}

func flatten(in interface{}, current_path string) ([]node, []link) {
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
			target := current_path + "-" + key

			nodes = append(nodes, node{
				name:  target,
				label: key,
			})

			if current_path != "" {
				links = append(links, link{
					left:  current_path,
					right: target,
				})
			}

			new_nodes, new_links := flatten(value, target)

			nodes = append(nodes, new_nodes...)
			links = append(links, new_links...)
		}

		return nodes, links
	case reflect.Array, reflect.Slice:
		for index := 0; index < val.Len(); index++ {
			value := val.Index(index).Interface()

			target := current_path + "-" + fmt.Sprint(index)

			nodes = append(nodes, node{
				name:  target,
				label: fmt.Sprintf("[%d]", index),
			})

			if current_path != "" {
				links = append(links, link{
					left:  current_path,
					right: target,
				})
			}

			new_nodes, new_links := flatten(value, target)

			nodes = append(nodes, new_nodes...)
			links = append(links, new_links...)
		}

		return nodes, links
	default:
		target := current_path + "=content"

		nodes = append(nodes, node{
			name:  target,
			label: fmt.Sprint(in),
		})

		links = append(links, link{
			left:  current_path,
			right: target,
		})

		return nodes, links
	}
}
