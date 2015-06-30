package parser

import (
	"reflect"
	"testing"
)

func TestXml(t *testing.T) {
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

	actual, err := Xml([]byte(input))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected response:\n%#v\nvs\n%#v", actual, expected)
	}
}

func TestXmlBadInput(t *testing.T) {
	input := `<things>
		<item id="1">Foo</item>
		</item>
	</things>`

	_, err := Xml([]byte(input))

	if err == nil {
		t.Errorf("Expected error")
	}
}
