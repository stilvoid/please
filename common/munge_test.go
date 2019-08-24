package common

import (
	"reflect"
	"testing"
)

type mungeCase struct {
	left, right, expected interface{}
}

func TestSuccess(t *testing.T) {
	cases := []mungeCase{
		// Matching types
		{"foo", "bar", "bar"},
		{17, 19, 19},
		{0.1, 7.9, 7.9},
		{true, false, false},

		// Mismatched types
		{"foo", 17, 17},
		{"bar", []string{"baz"}, []string{"baz"}},
		{[]string{"quux"}, "mooz", "mooz"},
		{[]string{"xyzzy"}, map[string]interface{}{}, map[string]interface{}{}},

		// Slices
		{
			[]string{},
			[]string{"foo"},
			[]interface{}{"foo"},
		},
		{
			[]string{"foo"},
			[]string{"foo", "bar"},
			[]interface{}{"foo", "bar"},
		},
		{
			[]string{"foo", "bar"},
			[]string{"baz"},
			[]interface{}{"baz", "bar"},
		},
		{
			[]interface{}{"foo", []string{"bar", "baz"}},
			[]interface{}{"quux", []string{"mooz"}},
			[]interface{}{"quux", []interface{}{"mooz", "baz"}},
		},

		// Maps
		{
			map[string]string{},
			map[string]string{"foo": "bar"},
			map[interface{}]interface{}{"foo": "bar"},
		},
		{
			map[string]string{"foo": "bar"},
			map[string]string{},
			map[interface{}]interface{}{"foo": "bar"},
		},
		{
			map[string]string{"foo": "bar"},
			map[string]interface{}{"foo": map[string]string{"bar": "baz"}},
			map[interface{}]interface{}{"foo": map[string]string{"bar": "baz"}},
		},
	}

	for _, testCase := range cases {
		actual := Munge(testCase.left, testCase.right)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Falsely munged %T '%v' and %T '%v' into %T '%v'",
				testCase.left, testCase.left,
				testCase.right, testCase.right,
				actual, actual,
			)
		}
	}
}
