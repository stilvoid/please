package formatters

import "testing"

func TestXML(t *testing.T) {
	expecteds := []string{
		"<doc>123</doc>",
		"<doc>456.789</doc>",
		"<doc>abc</doc>",
		"<doc>true</doc>",
		"<doc>false</doc>",
		"<doc></doc>",
		"<doc>\n  <element>123</element>\n  <element>abc</element>\n</doc>",
		"<foo>bar</foo>",
		"<doc>\n  <123>baz</123>\n  <123>quux</123>\n</doc>",
		"<true>\n<null/>\n</true>",
		"<doc>\n  <element>456</element>\n  <element>def</element>\n  <3>4</3>\n</doc>",
	}

	if len(expecteds) != len(testCases) {
		//t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatXML(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
