package parser

import (
	"testing"
)

func TestFilter(t *testing.T) {
	input := map[string]interface{}{
		"top": "I am the top",
		"bottom": map[int]interface{}{
			16: "left",
			13: []string{
				"right 1",
				"right 2",
			},
		},
	}

	cases := map[string]string{
		"top":         "I am the top",
		"bottom.16":   "left",
		"bottom.13.1": "right 2",
	}

	for path, expected := range cases {
		actual, err := Filter(input, path)

		if err != nil {
			t.Fail()
		}

		if expected != actual {
			t.Errorf("Case failed: %v vs %v", expected, actual)
		}
	}
}

func TestFilterBadKey(t *testing.T) {
	input := map[string]interface{}{
		"foo": "bar",
	}

	val, err := Filter(input, "not foo")

	if err == nil || err.Error() != "Key does not exist: not foo" {
		t.Errorf("Unexpected return values: %v, %v", val, err)
	}
}
