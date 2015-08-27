package util

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestFilter(t *testing.T) {
	input := map[string]interface{}{
		"top": "I am the top",
		"bottom": map[int]interface{}{
			16: "left",
			13: []interface{}{
				"right 1",
				map[string]interface{}{
					"deeper": "we go",
				},
			},
		},
	}

	cases := map[string]interface{}{
		"top":         "I am the top",
		"bottom":      input["bottom"],
		"bottom.16":   input["bottom"].(map[int]interface{})[16],
		"bottom.13":   input["bottom"].(map[int]interface{})[13],
		"bottom.13.0": input["bottom"].(map[int]interface{})[13].([]interface{})[0],
		"bottom.13.1": input["bottom"].(map[int]interface{})[13].([]interface{})[1],
	}

	for path, expected := range cases {
		actual, err := Filter(input, path)

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("case failed: %v vs %v", expected, actual)
		}
	}
}

func TestStarFilterSlice(t *testing.T) {
	input := map[string]interface{}{
		"top": []string{
			"foo",
			"bar",
			"baz",
		},
		"bottom": []interface{}{
			map[string]int{
				"id": 1,
			},
			map[string]int{
				"id": 2,
			},
			map[string]int{
				"id": 3,
			},
		},
	}

	cases := map[string]interface{}{
		"top.*":       []interface{}{"foo", "bar", "baz"},
		"bottom.*.id": []interface{}{1, 2, 3},
	}

	for path, expected := range cases {
		actual, err := Filter(input, path)

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("case failed: %v vs %v", expected, actual)
		}
	}
}

func TestStarFilterMap(t *testing.T) {
	input := map[string]interface{}{
		"top": map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		},
		"bottom": map[string]interface{}{
			"a": map[string]int{
				"id": 4,
			},
			"b": map[string]int{
				"id": 5,
			},
			"c": map[string]int{
				"id": 6,
			},
		},
	}

	cases := map[string]interface{}{
		"top.*":       []int{1, 2, 3},
		"bottom.*.id": []int{4, 5, 6},
	}

	for path, expected := range cases {
		actual, err := Filter(input, path)

		val := reflect.ValueOf(actual)

		// Some "fun" type conversion here to make sure things work ok
		intActual := make([]int, val.Len())

		for i := 0; i < val.Len(); i++ {
			fmt.Println(val.Index(i))
			intActual[i] = val.Index(i).Interface().(int)
		}

		sort.Ints(intActual)

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(expected, intActual) {
			t.Errorf("case failed: %v vs %v", expected, actual)
		}
	}
}

func TestFilterBadKey(t *testing.T) {
	input := map[string]interface{}{
		"foo": "bar",
	}

	val, err := Filter(input, "not foo")

	if err == nil || err.Error() != "key does not exist: not foo" {
		t.Errorf("unexpected return values: %v, %v", val, err)
	}
}

func TestStarFilterBadKey(t *testing.T) {
	input := []interface{}{
		map[string]int{
			"id": 1,
		},
		map[string]int{
			"id": 2,
		},
		map[string]int{
			"notid": 3,
		},
	}

	val, err := Filter(input, "*.id")

	if err == nil || err.Error() != "key does not exist: id" {
		t.Errorf("unexpected return values: %v, %v", val, err)
	}
}
