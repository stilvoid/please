package parsers

import (
	"reflect"
	"testing"
)

func TestXML(t *testing.T) {
	input := `<things>
		<item id="1">Foo</item>
		<item id="2">
			Bar
			<child>Rock</child>
		</item>
	</things>`

	expected := map[string]interface{}{
		"things": map[string]interface{}{
			"item": []interface{}{
				map[string]interface{}{
					"-id":   "1",
					"#text": "Foo",
				},
				map[string]interface{}{
					"-id":   "2",
					"#text": "Bar",
					"child": "Rock",
				},
			},
		},
	}

	actual, err := parseXML([]byte(input))

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

	_, err := parseXML([]byte(input))

	if err == nil {
		t.Errorf("expected error")
	}
}
