package parse_test

import (
	"testing"

	"github.com/stilvoid/please/parse"
)

const jsonInput = `{
    "description": "some example json",
    "an array": [
        "first entry",
        {
            "nested": "object"
        },
        ["nested", "array"]
    ],
    "child": {
        "with": "value"
    }
}`

const yamlInput = `an array:
- first entry
- nested: object
- - nested
  - array
child:
  with: value
description: some example json`

const xmlInput = `<doc>
  <description>some example json</description>
  <an_array>first entry</an_array>
  <an_array>
    <nested>object</nested>
  </an_array>
    <an_array>nested</an_array>
    <an_array>array</an_array>
  <child>
    <with>value</with>
  </child>
</doc>`

const queryInput = `an+array=0%3Dfirst%2Bentry%261%3Dnested%253Dobject%262%3D0%253Dnested%25261%253Darray&child=with%3Dvalue&description=some+example+json`

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse.Json([]byte(jsonInput))
	}
}

func BenchmarkYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse.Yaml([]byte(yamlInput))
	}
}

func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse.Xml([]byte(xmlInput))
	}
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse.Query([]byte(queryInput))
	}
}

func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse.Html([]byte(xmlInput))
	}
}
