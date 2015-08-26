package parsers

import (
	"reflect"
	"testing"
)

func TestCSV(t *testing.T) {
	input := "col1,col2\n" +
		"\"1-1\",\"1-2\"\n" +
		"2-1,2-2"

	expected := [][]string{
		[]string{"col1", "col2"},
		[]string{"1-1", "1-2"},
		[]string{"2-1", "2-2"},
	}

	actual, err := parseCSV([]byte(input))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
