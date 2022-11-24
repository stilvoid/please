package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestQuery(t *testing.T) {
	cases := map[string]map[string][]string{
		"foo=bar": {
			"foo": {"bar"},
		},
		"foo=bar&foo=baz": {
			"foo": {"bar", "baz"},
		},
		"foo=bar&baz=quux": {
			"foo": {"bar"},
			"baz": {"quux"},
		},
	}

	for input, expected := range cases {
		actual, err := parse.Query([]byte(input))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%#v', want '%#v'", actual, expected)
		}
	}
}
