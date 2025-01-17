package format_test

import (
	"testing"

	"github.com/stilvoid/please/format"
)

var input = map[any]any{
	"description": "some example json",
	"an array": []any{
		"first entry",
		map[any]any{
			"nested": "object",
		},
		[]any{"nested", "array"},
	},
	"child": map[any]any{
		"with": "value",
	},
}

func BenchmarkBash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Bash(input)
	}
}

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Dot(input)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Json(input)
	}
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Query(input)
	}
}

func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Xml(input)
	}
}

func BenchmarkYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format.Yaml(input)
	}
}
