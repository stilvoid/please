package formatters

import "testing"

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
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatDot(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
