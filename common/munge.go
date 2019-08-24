package common

import (
	"math"
	"reflect"
)

func Munge(left, right interface{}) interface{} {
	return munge(reflect.ValueOf(left), reflect.ValueOf(right)).Interface()
}

func munge(left, right reflect.Value) reflect.Value {
	// Reload in case we've been previously munged
	left = reflect.ValueOf(left.Interface())
	right = reflect.ValueOf(right.Interface())

	if left.Kind() == right.Kind() {
		switch left.Kind() {
		case reflect.Slice, reflect.Array:
			max := int(math.Max(float64(left.Len()), float64(right.Len())))

			out := reflect.ValueOf(make([]interface{}, max, max))

			for i := 0; i < max; i++ {
				if i >= left.Len() {
					out.Index(i).Set(right.Index(i))
				} else if i >= right.Len() {
					out.Index(i).Set(left.Index(i))
				} else {
					out.Index(i).Set(munge(left.Index(i), right.Index(i)))
				}
			}

			return out
		case reflect.Map:
			out := reflect.ValueOf(make(map[interface{}]interface{}))

			for _, key := range left.MapKeys() {
				out.SetMapIndex(key, left.MapIndex(key))
			}

			for _, key := range right.MapKeys() {
				l := left.MapIndex(key)
				r := right.MapIndex(key)

				if !l.IsValid() {
					out.SetMapIndex(key, r)
				} else {
					out.SetMapIndex(key, munge(l, r))
				}
			}

			return out
		}
	}

	return right
}
