package parser

import (
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	inputs := []string{
		`{"foo": "bar", "baz": 123, "quux": [1,2,[3,4],{"a": false}]}`,
		`"I am a fish"`,
		`["hello", 123, {"b": ["cake", true]}]`,
		`[1,2,3,4,5]`,
	}

	expecteds := []interface{}{
		map[string]interface{}{
			"foo": "bar",
			"baz": 123.0,
			"quux": []interface{}{
				1.0,
				2.0,
				[]interface{}{
					3.0,
					4.0,
				},
				map[string]interface{}{
					"a": false,
				},
			},
		},
		"I am a fish",
		[]interface{}{
			"hello",
			123.0,
			map[string]interface{}{
				"b": []interface{}{
					"cake",
					true,
				},
			},
		},
		[]interface{}{1.0, 2.0, 3.0, 4.0, 5.0},
	}

	for i := range inputs {
		input := inputs[i]
		expected := expecteds[i]

		actual, err := parseJson([]byte(input))

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected response:\n%#v\nvs\n%#v", actual, expected)
		}
	}
}
