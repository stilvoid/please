package parse_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/please/parse"
)

func TestToml(t *testing.T) {
	inputs := []string{
		`foo="bar"
baz=123
quux=[1,2,[3,4],{a=false}]`,
		`test=["hello",123,{b=["cake",true]}]`,
		`[test]
cake=[1,2,3,4,5]`,
	}

	expecteds := []any{
		map[string]any{
			"foo": "bar",
			"baz": int64(123),
			"quux": []any{
				int64(1),
				int64(2),
				[]any{
					int64(3),
					int64(4),
				},
				map[string]any{
					"a": false,
				},
			},
		},
		map[string]any{
			"test": []any{
				"hello",
				int64(123),
				map[string]any{
					"b": []any{
						"cake",
						true,
					},
				},
			},
		},
		map[string]any{
			"test": map[string]any{
				"cake": []any{int64(1), int64(2), int64(3), int64(4), int64(5)},
			},
		},
	}

	for i := range inputs {
		input := inputs[i]
		expected := expecteds[i]

		actual, err := parse.Toml([]byte(input))

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}
