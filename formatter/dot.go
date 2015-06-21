package formatter

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

func Dot(in interface{}, path string) (out string) {
	nodes, links := flatten(in, path, "root")

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

func flatten(in interface{}, path string, current_path string) ([]node, []link) {
	var nodes []node
	var links []link

	if in == nil {
		return nodes, links
	}

	val := reflect.ValueOf(in)

	split_path := strings.SplitN(path, ".", 2)

	this_path := split_path[0]
	var next_path string

	if len(split_path) > 1 {
		next_path = split_path[1]
	}

	switch val.Kind() {
	case reflect.Map:
		vv := in.(map[string]interface{})

		if this_path != "" {
			if _, ok := vv[this_path]; !ok {
				fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
				os.Exit(1)
			}

			return flatten(vv[this_path], next_path, this_path)
		}

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

			new_nodes, new_links := flatten(value, path, target)

			nodes = append(nodes, new_nodes...)
			links = append(links, new_links...)
		}

		return nodes, links
	case reflect.Array, reflect.Slice:
		if this_path != "" {
			index, err := strconv.Atoi(this_path)

			if err != nil || index < 0 || index >= val.Len() {
				fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
				os.Exit(1)
			}

			return flatten(val.Index(index).Interface(), next_path, this_path)
		}

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

			new_nodes, new_links := flatten(value, path, target)

			nodes = append(nodes, new_nodes...)
			links = append(links, new_links...)
		}

		return nodes, links
	default:
		if this_path != "" {
			fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
			os.Exit(1)
		}

		links = append(links, link{
			left:  current_path,
			right: fmt.Sprint(in),
		})

		return nodes, links
	}
}
