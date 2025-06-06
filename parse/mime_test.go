package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestMIME(t *testing.T) {
	input := "Test:    123\n" +
		"Multiple-header: value1\n" +
		"multiple-Header: Value2\n" +
		"\n" +
		"This is the body."

	expected := map[string]any{
		"headers": map[string]any{
			"Test":            []string{"123"},
			"Multiple-Header": []string{"value1", "Value2"},
		},
		"body": "This is the body.",
	}

	actual, err := parse.Mime([]byte(input))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
