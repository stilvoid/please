package parser

import (
	"reflect"
	"testing"
)

func TestMime(t *testing.T) {
	input := "Test:    123\n" +
		"Multiple-header: value1\n" +
		"multiple-Header: Value2\n" +
		"\n" +
		"This is the body."

	expected := map[string]interface{}{
		"headers": map[string]interface{}{
			"Test":            []string{"123"},
			"Multiple-Header": []string{"value1", "Value2"},
		},
		"body": "This is the body.",
	}

	actual, err := parseMime([]byte(input))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
