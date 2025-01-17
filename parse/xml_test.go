package parse_test

import (
	"reflect"
	"testing"

	"github.com/stilvoid/please/parse"
)

func TestXML(t *testing.T) {
	input := `<things>
		<item id="1">Foo</item>
		<item id="2">
			Bar
			<child>Rock</child>
		</item>
	</things>`

	expected := map[string]any{
		"things": map[string]any{
			"item": []any{
				map[string]any{
					"-id":   "1",
					"#text": "Foo",
				},
				map[string]any{
					"-id":   "2",
					"#text": "Bar",
					"child": "Rock",
				},
			},
		},
	}

	actual, err := parse.Xml([]byte(input))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}

func TestXMLBadInput(t *testing.T) {
	input := `<things>
		<item id="1">Foo</item>
		</item>
	</things>`

	_, err := parse.Xml([]byte(input))

	if err == nil {
		t.Errorf("expected error")
	}
}
