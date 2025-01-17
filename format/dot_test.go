package format_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/please/format"
)

func TestDot(t *testing.T) {
	expecteds := []string{
		`graph{
"root" [label="123"];
}`,

		`graph{
"root" [label="456.789"];
}`,

		`graph{
"root" [label="abc"];
}`,

		`graph{
"root" [label="true"];
}`,

		`graph{
"root" [label="false"];
}`,

		`graph{
"root" [label="<nil>"];
}`,

		`graph{
"root" [label="[array]"];
"root" -- "root-array-0";
"root-array-0" [label="123"];
"root" -- "root-array-1";
"root-array-1" [label="abc"];
}`,

		`graph{
"root" [label="[map]"];
"root" -- "root-map-0";
"root-map-0" [label="foo"];
"root-map-0" -- "root-map-0=content";
"root-map-0=content" [label="bar"];
}`,

		`graph{
"root" [label="[map]"];
"root" -- "root-map-0";
"root-map-0" [label="123"];
"root-map-0" -- "root-map-0=content";
"root-map-0=content" [label="[array]"];
"root-map-0=content" -- "root-map-0=content-array-0";
"root-map-0=content-array-0" [label="baz"];
"root-map-0=content" -- "root-map-0=content-array-1";
"root-map-0=content-array-1" [label="quux"];
}`,

		`graph{
"root" [label="[map]"];
"root" -- "root-map-0";
"root-map-0" [label="true"];
"root-map-0" -- "root-map-0=content";
"root-map-0=content" [label="[map]"];
"root-map-0=content" -- "root-map-0=content-map-0";
"root-map-0=content-map-0" [label="null"];
"root-map-0=content-map-0" -- "root-map-0=content-map-0=content";
"root-map-0=content-map-0=content" [label="<nil>"];
}`,

		`graph{
"root" [label="[array]"];
"root" -- "root-array-0";
"root-array-0" [label="456"];
"root" -- "root-array-1";
"root-array-1" [label="def"];
"root" -- "root-array-2";
"root-array-2" [label="[map]"];
"root-array-2" -- "root-array-2-map-0";
"root-array-2-map-0" [label="3"];
"root-array-2-map-0" -- "root-array-2-map-0=content";
"root-array-2-map-0=content" [label="4"];
"root" -- "root-array-3";
"root-array-3" [label="[array]"];
"root-array-3" -- "root-array-3-array-0";
"root-array-3-array-0" [label="first"];
"root-array-3" -- "root-array-3-array-1";
"root-array-3-array-1" [label="second"];
"root-array-3" -- "root-array-3-array-2";
"root-array-3-array-2" [label="[array]"];
"root-array-3-array-2" -- "root-array-3-array-2-array-0";
"root-array-3-array-2-array-0" [label="deeper"];
}`,

		`graph{
"root" [label="[map]"];
"root" -- "root-map-0";
"root-map-0" [label="Array"];
"root-map-0" -- "root-map-0=content";
"root-map-0=content" [label="[array]"];
"root-map-0=content" -- "root-map-0=content-array-0";
"root-map-0=content-array-0" [label="def"];
"root-map-0=content" -- "root-map-0=content-array-1";
"root-map-0=content-array-1" [label="456"];
"root-map-0=content" -- "root-map-0=content-array-2";
"root-map-0=content-array-2" [label="true"];
"root-map-0=content" -- "root-map-0=content-array-3";
"root-map-0=content-array-3" [label="false"];
"root-map-0=content" -- "root-map-0=content-array-4";
"root-map-0=content-array-4" [label="<nil>"];
"root" -- "root-map-1";
"root-map-1" [label="Map"];
"root-map-1" -- "root-map-1=content";
"root-map-1=content" [label="[map]"];
"root-map-1=content" -- "root-map-1=content-map-0";
"root-map-1=content-map-0" [label="456"];
"root-map-1=content-map-0" -- "root-map-1=content-map-0=content";
"root-map-1=content-map-0=content" [label="def"];
"root-map-1=content" -- "root-map-1=content-map-1";
"root-map-1=content-map-1" [label="foo"];
"root-map-1=content-map-1" -- "root-map-1=content-map-1=content";
"root-map-1=content-map-1=content" [label="123"];
"root" -- "root-map-2";
"root-map-2" [label="Name"];
"root-map-2" -- "root-map-2=content";
"root-map-2=content" [label="abc"];
"root" -- "root-map-3";
"root-map-3" [label="Number"];
"root-map-3" -- "root-map-3=content";
"root-map-3=content" [label="(12+3i)"];
}`,
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := format.Dot(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}
