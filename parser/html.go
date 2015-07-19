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

func Html(input []byte) (interface{}, error) {
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
			existing_value, exists := out[c.Data]

			new_value := formatHtml(c)

			if !exists {
				out[c.Data] = new_value
			} else {
				val := reflect.ValueOf(existing_value)

				kind := val.Kind()

				if kind != reflect.Array && kind != reflect.Slice {
					out[c.Data] = []interface{}{
						existing_value,
						new_value,
					}
				} else {
					// *this* is sick
					out[c.Data] = reflect.Append(val, reflect.ValueOf(new_value)).Interface().([]interface{})
				}
			}
		}
	}

	return out
}

func init() {
	Parsers["html"] = parser{
		parse:   Html,
		prefers: []string{"xml", "yaml"},
	}
}
