package parser

import (
	"reflect"
	"testing"
)

func TestHtml(t *testing.T) {
	input := `<html>
		<head>
			<title>Test</title>
		</head>
		
		<body id="mybody">
			<p>Hello</p>
		</body>
	</html>`

	expected := map[string]interface{}{
		"html": map[string]interface{}{
			"head": map[string]interface{}{
				"title": map[string]interface{}{
					"#text": "Test",
				},
			},
			"body": map[string]interface{}{
				"-id": "mybody",
				"p": map[string]interface{}{
					"#text": "Hello",
				},
			},
		},
	}

	actual, err := Html([]byte(input))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
