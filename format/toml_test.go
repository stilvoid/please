package format_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/please/format"
)

func TestToml(t *testing.T) {
	expecteds := []string{
		`123`,
		`456.789`,
		`"abc"`,
		`true`,
		`false`,
		`"null"`,
		`[123, "abc"]`,
		`foo = "bar"`,
		`123 = ["baz", "quux"]`,
		`[true]
  null = "null"`,
		`[456, "def", {3 = 4}, ["first", "second", ["deeper"]]]`,
		`Array = ["def", 456, true, false, "null"]
Name = "abc"
Number = "(12+3i)"

[Map]
  456 = "def"
  foo = 123`,
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := format.Toml(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}
