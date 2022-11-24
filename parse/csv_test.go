package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestCSV(t *testing.T) {
	input := "col1,col2\n" +
		"\"1-1\",\"1-2\"\n" +
		"2-1,2-2"

	expected := [][]string{
		{"col1", "col2"},
		{"1-1", "1-2"},
		{"2-1", "2-2"},
	}

	actual, err := parse.Csv([]byte(input))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
