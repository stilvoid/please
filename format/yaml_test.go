package format

import "testing"

func TestYAML(t *testing.T) {
	expecteds := []string{
		"123",
		"456.789",
		"abc",
		"true",
		"false",
		"null",
		"- 123\n- abc",
		"foo: bar",
		"123:\n- baz\n- quux",
		"true:\n  null: null",
		"- 456\n- def\n- 3: 4\n- - first\n  - second\n  - - deeper",
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatYAML(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected:\n'%v'\nwant:\n'%v'", actual, expected)
		}
	}
}
