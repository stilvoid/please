package parser

import (
	"reflect"
	"testing"
)

func TestCsv(t *testing.T) {
	input := "col1,col2\n" +
		"\"1-1\",\"1-2\"\n" +
		"2-1,2-2"

	expected := [][]string{
		[]string{"col1", "col2"},
		[]string{"1-1", "1-2"},
		[]string{"2-1", "2-2"},
	}

	actual, err := Csv([]byte(input))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}
