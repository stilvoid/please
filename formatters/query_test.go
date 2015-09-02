package formatters

import "testing"

func TestQuery(t *testing.T) {
	expecteds := []string{
		"123",
		"456.789",
		"abc",
		"true",
		"false",
		"",
		"0=123&1=abc",
		"foo=bar",
		"123=0%3Dbaz%261%3Dquux",
		"true=null%3D",
		"0=456&1=def&2=3%3D4",
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatQuery(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}