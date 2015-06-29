package util

import (
	"reflect"
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

func TestForceStringKeys(t *testing.T) {
	input := map[int]interface{}{
		13: []interface{}{
			"foo",
			map[int]interface{}{
				0:   "none",
				100: "some",
			},
		},
		66: map[int]interface{}{
			1: []interface{}{"foo", "bar"},
			2: "two",
		},
	}

	expected := map[string]interface{}{
		"13": []interface{}{
			"foo",
			map[string]interface{}{
				"0":   "none",
				"100": "some",
			},
		},
		"66": map[string]interface{}{
			"1": []interface{}{"foo", "bar"},
			"2": "two",
		},
	}

	actual := ForceStringKeys(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result:\n%#v\nvs\n%#v\n", actual, expected)
	}
}
