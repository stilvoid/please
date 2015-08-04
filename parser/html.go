package parser

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

type node struct {
	Node  interface{}   `xml:",any"`
	List  []interface{} `xml:",any"`
	Value interface{}   `xml:",any"`
}

func parseHTML(input []byte) (interface{}, error) {
	var parsed interface{}

	doc, err := html.Parse(bytes.NewReader(input))

	if err == nil {
		parsed = formatHTML(doc)
	}

	return parsed, err
}

func formatHTML(n *html.Node) map[string]interface{} {
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

			newValue := formatHTML(c)

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
		parse:   parseHTML,
		prefers: []string{"xml", "yaml"},
	}
}
