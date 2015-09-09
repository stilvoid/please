package formatters

import "testing"

var input = map[interface{}]interface{}{
	"description": "some example json",
	"an array": []interface{}{
		"first entry",
		map[interface{}]interface{}{
			"nested": "object",
		},
		[]interface{}{"nested", "array"},
	},
	"child": map[interface{}]interface{}{
		"with": "value",
	},
}

func BenchmarkBash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatBash(input)
	}
}

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatDot(input)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatJSON(input)
	}
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatQuery(input)
	}
}

func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatXML(input)
	}
}

func BenchmarkYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatYAML(input)
	}
}
