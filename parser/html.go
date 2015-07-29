package parser

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"reflect"
	"strings"
)

type node struct {
	node  interface{}   `xml:",any"`
	list  []interface{} `xml:",any"`
	value interface{}   `xml:",any"`
}

func parseHtml(input []byte) (interface{}, error) {
	var parsed interface{}

	doc, err := html.Parse(bytes.NewReader(input))

	if err == nil {
		parsed = formatHtml(doc)
	}

	return parsed, err
}

func formatHtml(n *html.Node) map[string]interface{} {
	out := make(map[string]interface{})

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

			newValue := formatHtml(c)

			if !exists {
				out[c.Data] = newValue
			} else {
				val := reflect.ValueOf(existingValue)

				kind := val.Kind()

				if kind != reflect.Array && kind != reflect.Slice {
					out[c.Data] = []interface{}{
						existingValue,
						newValue,
					}
				} else {
					// *this* is sick
					out[c.Data] = reflect.Append(val, reflect.ValueOf(newValue)).Interface().([]interface{})
				}
			}
		}
	}

	return out
}

func init() {
	parsers["html"] = parser{
		parse:   parseHtml,
		prefers: []string{"xml", "yaml"},
	}
}
