package util

import (
	"testing"
)

func TestSortKeys(t *testing.T) {
	input := map[interface{}]interface{}{
		"2": 0,
		0:   1, // We should be converting non-strings
		"1": 2,
	}

	expected := []string{"0", "1", "2"}

	actual := SortKeys(input)

	if len(actual) != len(expected) {
		t.Errorf("Lengths differ: %d vs %d", len(actual), len(expected))
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Values differ: %v vs %v", actual, expected)
		}
	}
}

func TestSortKeysNotMap(t *testing.T) {
	defer func() {
		err := recover()

		if err != "Not a map!" {
			t.Errorf("Didn't fail correctly: %v", err)
		}
	}()

	input := "I am not a map"

	SortKeys(input)
}

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
		actual := Filter(input, path)

		if expected != actual {
			t.Errorf("Case failed: %v vs %v", expected, actual)
		}
	}
}
