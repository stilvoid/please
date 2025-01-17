package format_test

var testCases = []interface{}{
	123,
	456.789,
	"abc",
	true,
	false,
	nil,
	[]interface{}{123, "abc"},
	map[interface{}]interface{}{ // A map
		"foo": "bar", // To a value
	},
	map[interface{}]interface{}{ // A map
		123: []interface{}{"baz", "quux"}, // To an array
	},
	map[interface{}]interface{}{ // A map
		true: map[interface{}]interface{}{ // To another map
			nil: nil,
		},
	},
	[]interface{}{ // An array
		456, // Of values
		"def",
		map[interface{}]interface{}{ // With a map
			3: 4,
		},
		[]interface{}{ // And another array
			"first",
			"second",
			[]interface{}{ // Triply embedded array
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
