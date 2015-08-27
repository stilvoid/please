package formatters

import "testing"

func TestJSON(t *testing.T) {
	expecteds := []string{
		"123",
		"456.789",
		"\"abc\"",
		"true",
		"false",
		"null",
		"[\n  123,\n  \"abc\"\n]",
		"{\n  \"foo\": \"bar\"\n}",
		"{\n  \"123\": [\n    \"baz\",\n    \"quux\"\n  ]\n}",
		"{\n  \"true\": {\n    \"null\": null\n  }\n}",
		"[\n  456,\n  \"def\",\n  {\n    \"3\": 4\n  }\n]",
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatJSON(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
