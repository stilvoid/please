package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestJSON(t *testing.T) {
	inputs := []string{
		`{"foo": "bar", "baz": 123, "quux": [1,2,[3,4],{"a": false}]}`,
		`"I am a fish"`,
		`["hello", 123, {"b": ["cake", true]}]`,
		`[1,2,3,4,5]`,
	}

	expecteds := []any{
		map[string]any{
			"foo": "bar",
			"baz": 123.0,
			"quux": []any{
				1.0,
				2.0,
				[]any{
					3.0,
					4.0,
				},
				map[string]any{
					"a": false,
				},
			},
		},
		"I am a fish",
		[]any{
			"hello",
			123.0,
			map[string]any{
				"b": []any{
					"cake",
					true,
				},
			},
		},
		[]any{1.0, 2.0, 3.0, 4.0, 5.0},
	}

	for i := range inputs {
		input := inputs[i]
		expected := expecteds[i]

		actual, err := parse.Json([]byte(input))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
		}
	}
}
