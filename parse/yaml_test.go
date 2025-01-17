package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestYAML(t *testing.T) {
	inputs := []string{
		"foo: bar\nbaz: 123\nquux:\n  - 1\n  - 2\n  - - 3\n    - 4\n  - {a: false}",
		"I am a fish",
		"- hello\n- \"123\"\n- b:\n  - cake\n  - true",
		"[1,2,3,4,5]",
	}

	expecteds := []any{
		map[any]any{
			"foo": "bar",
			"baz": 123,
			"quux": []any{
				1,
				2,
				[]any{
					3,
					4,
				},
				map[any]any{
					"a": false,
				},
			},
		},
		"I am a fish",
		[]any{
			"hello",
			"123",
			map[any]any{
				"b": []any{
					"cake",
					true,
				},
			},
		},
		[]any{1, 2, 3, 4, 5},
	}

	for i := range inputs {
		input := inputs[i]
		expected := expecteds[i]

		actual, err := parse.Yaml([]byte(input))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
		}
	}
}
