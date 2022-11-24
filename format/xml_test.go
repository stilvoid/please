package format

import "testing"

func TestXML(t *testing.T) {
	xmlTestCases := append(testCases, map[interface{}]interface{}{ // XML style
		"top": map[interface{}]interface{}{
			"#text":      []interface{}{"Some text", "more text"},
			"-attribute": "value",
			"child": map[interface{}]interface{}{
				"#text": "child text",
				"-type": "clever",
			},
		},
	})

	expecteds := []string{
		"<root>123</root>",
		"<root>456.789</root>",
		"<root>abc</root>",
		"<root>true</root>",
		"<root>false</root>",
		"<root>&lt;nil&gt;</root>",
		`<root>
  <tag>123</tag>
  <tag>abc</tag>
</root>`,
		`<root>
  <foo>bar</foo>
</root>`,
		`<root>
  <_23>baz</_23>
  <_23>quux</_23>
</root>`,
		`<root>
  <true>
    <null>&lt;nil&gt;</null>
  </true>
</root>`,
		`<root>
  <tag>456</tag>
  <tag>def</tag>
  <tag>
    <_>4</_>
  </tag>
  <tag>
    <tag>first</tag>
    <tag>second</tag>
    <tag>
      <tag>deeper</tag>
    </tag>
  </tag>
</root>`,
		`<root>
  <top attribute="value">
    Some text
    more text
    <child type="clever">child text</child>
  </top>
</root>`,
	}

	if len(expecteds) != len(xmlTestCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := xmlTestCases[i]

		actual, err := formatXML(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected:\n'%v'\nwant:\n'%v'", actual, expected)
		}
	}
}
