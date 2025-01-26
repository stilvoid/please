package parse

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

func Html(input []byte) (any, error) {
	var parsed any

	doc, err := html.Parse(bytes.NewReader(input))

	if err == nil {
		parsed = formatHTML(doc)
	}

	return parsed, err
}

func formatHTML(n *html.Node) map[string]any {
	out := make(map[string]any)

	for _, a := range n.Attr {
		out[fmt.Sprintf("-%s", a.Key)] = a.Val
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text := strings.TrimSpace(c.Data)

			if text != "" {
				out["#text"] = c.Data
			}
		} else {
			existingValue, exists := out[c.Data]

			newValue := formatHTML(c)

			if !exists {
				out[c.Data] = newValue
			} else {
				val := reflect.ValueOf(existingValue)

				kind := val.Kind()

				if kind != reflect.Array && kind != reflect.Slice {
					out[c.Data] = []any{
						existingValue,
						newValue,
					}
				} else {
					// *this* is sick
					out[c.Data] = reflect.Append(val, reflect.ValueOf(newValue)).Interface().([]any)
				}
			}
		}
	}

	return out
}
