package format_test

var testCases = []any{
	123,
	456.789,
	"abc",
	true,
	false,
	nil,
	[]any{123, "abc"},
	map[any]any{ // A map
		"foo": "bar", // To a value
	},
	map[any]any{ // A map
		123: []any{"baz", "quux"}, // To an array
	},
	map[any]any{ // A map
		true: map[any]any{ // To another map
			nil: nil,
		},
	},
	[]any{ // An array
		456, // Of values
		"def",
		map[any]any{ // With a map
			3: 4,
		},
		[]any{ // And another array
			"first",
			"second",
			[]any{ // Triply embedded array
				"deeper",
			},
		},
	},
	struct {
		Name   string
		Number complex64
		Array  []any
		Map    map[any]any
	}{
		Name:   "abc",
		Number: complex(12, 3),
		Array: []any{
			"def",
			456,
			true,
			false,
			nil,
		},
		Map: map[any]any{
			"foo": 123,
			456:   "def",
		},
	},
}
