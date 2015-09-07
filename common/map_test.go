package common

import (
	"reflect"
	"testing"
)

func TestForceStringKeys(t *testing.T) {
	input := map[interface{}]interface{}{
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
		nil: "derf",
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
		"null": "derf",
	}

	actual := ForceStringKeys(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result:\n%#v\nvs\n%#v\n", actual, expected)
	}
}

func TestArraysToMaps(t *testing.T) {
	input := []interface{}{
		[]interface{}{123, 456},
		map[string]interface{}{
			"foo": []interface{}{789},
		},
	}

	expected := map[interface{}]interface{}{
		0: map[interface{}]interface{}{
			0: 123,
			1: 456,
		},
		1: map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				0: 789,
			},
		},
	}

	actual := ArraysToMaps(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected:\n%#v\nwant:\n%#v", actual, expected)
	}
}

func TestSortedKeys(t *testing.T) {
	inputs := []map[interface{}]interface{}{
		{
			"foo":  "bar",
			"baz":  "quux",
			"mooz": "xyzzy",
		},
		{
			2: "two",
			1: "one",
			0: "zero",
		},
		{
			"up": 0,
			1:    1,
			true: 2,
			nil:  3,
		},
	}

	expecteds := [][]interface{}{
		{"baz", "foo", "mooz"},
		{0, 1, 2},
		{1, nil, true, "up"},
	}

	for i, input := range inputs {
		actual := SortedKeys(input)

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected:\n%#v\nwant:\n%#v", actual, expected)
		}
	}
}
