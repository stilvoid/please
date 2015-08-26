package formatters

import (
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	inputs := []interface{}{
		map[string]string{
			"foo": "bar",
		},
		map[string][]string{
			"foo": {"bar", "baz"},
		},
		map[string]int{
			"foo": 123,
		},
		map[int]int{
			123: 456,
		},
		map[string]string{
			"foo": "bar",
			"baz": "quux",
		},
	}

	expecteds := []string{
		"foo=bar",
		"foo=bar&foo=baz",
		"foo=123",
		"123=456",
		"baz=quux&foo=bar",
	}

	for i, input := range inputs {
		actual, err := formatQuery(input)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
