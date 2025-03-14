package format_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/please/format"
)

func TestBash(t *testing.T) {
	expecteds := []string{
		`123`,
		`456.789`,
		`abc`,
		`true`,
		`false`,
		``,
		`([0]="123" [1]="abc")`,
		`([foo]="bar")`,
		`([123]="([0]=\"baz\" [1]=\"quux\")")`,
		`([true]="([null]=\"\")")`,
		`([0]="456" [1]="def" [2]="([3]=\"4\")" [3]="([0]=\"first\" [1]=\"second\" [2]=\"([0]=\\\"deeper\\\")\")")`,
		`([Array]="([0]=\"def\" [1]=\"456\" [2]=\"true\" [3]=\"false\" [4]=\"\")" [Map]="([456]=\"def\" [foo]=\"123\")" [Name]="abc" [Number]="(12+3i)")`,
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := format.Bash(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}
