package parser

import (
	"reflect"
	"testing"
)

func TestYAML(t *testing.T) {
	inputs := []string{
		"foo: bar\nbaz: 123\nquux:\n  - 1\n  - 2\n  - - 3\n    - 4\n  - {a: false}",
		"I am a fish",
		"- hello\n- \"123\"\n- b:\n  - cake\n  - true",
		"[1,2,3,4,5]",
	}

	expecteds := []interface{}{
		map[interface{}]interface{}{
			"foo": "bar",
			"baz": 123,
			"quux": []interface{}{
				1,
				2,
				[]interface{}{
					3,
					4,
				},
				map[interface{}]interface{}{
					"a": false,
				},
			},
		},
		"I am a fish",
		[]interface{}{
			"hello",
			"123",
			map[interface{}]interface{}{
				"b": []interface{}{
					"cake",
					true,
				},
			},
		},
		[]interface{}{1, 2, 3, 4, 5},
	}

	for i := range inputs {
		input := inputs[i]
		expected := expecteds[i]

		actual, err := parseYAML([]byte(input))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
		}
	}
}
