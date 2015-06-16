package parser

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type node struct {
	node  interface{}   `xml:",any"`
	list  []interface{} `xml:",any"`
	value interface{}   `xml:",any"`
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
			// FIXME - Deal with multiples of the same node type
			out[c.Data] = formatHtml(c)
		}
	}

	return out
}

func Html(input []byte, path string) (interface{}, error) {
	var parsed interface{}

	doc, err := html.Parse(bytes.NewReader(input))

	if err == nil {
		parsed = formatHtml(doc)
	}

	return parsed, err
}
