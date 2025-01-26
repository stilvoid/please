package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestHTML(t *testing.T) {
	input := `<html>
		<head>
			<title>Test</title>
		</head>
		
		<body id="mybody">
			<p>Hello</p>
		</body>
	</html>`

	expected := map[string]any{
		"html": map[string]any{
			"head": map[string]any{
				"title": map[string]any{
					"#text": "Test",
				},
			},
			"body": map[string]any{
				"-id": "mybody",
				"p": map[string]any{
					"#text": "Hello",
				},
			},
		},
	}

	actual, err := parse.Html([]byte(input))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
