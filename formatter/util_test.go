package formatter

import (
	"reflect"
	"testing"
)

func TestForceStringKeys(t *testing.T) {
	input := map[int]interface{}{
		13: []interface{}{
			"foo",
			map[int]interface{}{
				0:   "none",
				100: "some",
			},
		},
		66: map[int]interface{}{
			1: []interface{}{"foo", "bar"},
			2: "two",
		},
	}

	expected := map[string]interface{}{
		"13": []interface{}{
			"foo",
			map[string]interface{}{
				"0":   "none",
				"100": "some",
			},
		},
		"66": map[string]interface{}{
			"1": []interface{}{"foo", "bar"},
			"2": "two",
		},
	}

	actual := forceStringKeys(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result:\n%#v\nvs\n%#v\n", actual, expected)
	}
}
