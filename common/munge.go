package common

import (
	"math"
	"reflect"
)

var interfaceType reflect.Type

func init() {
	interfaceType = reflect.ValueOf(make([]interface{}, 0)).Type().Elem()
}

type FilterFunc func(reflect.Value, reflect.Value)

func Munge(left, right interface{}) interface{} {
	return munge(reflect.ValueOf(left), reflect.ValueOf(right), nil).Interface()
}

func MungeWithFilter(left, right interface{}, f FilterFunc) interface{} {
	return munge(reflect.ValueOf(left), reflect.ValueOf(right), f).Interface()
}

func munge(left, right reflect.Value, f FilterFunc) reflect.Value {
	// Reload in case we've been previously munged
	left = reflect.ValueOf(left.Interface())
	right = reflect.ValueOf(right.Interface())

	if f != nil {
		f(left, right)
	}

	if left.Kind() == right.Kind() {
		switch left.Kind() {
		case reflect.Slice, reflect.Array:
			var out reflect.Value

			max := int(math.Max(float64(left.Len()), float64(right.Len())))

			if right.Type().Elem().AssignableTo(left.Type().Elem()) {
				out = reflect.MakeSlice(left.Type(), max, max)
			} else {
				out = reflect.ValueOf(make([]interface{}, max, max))
			}

			for i := 0; i < max; i++ {
				if i >= left.Len() {
					out.Index(i).Set(right.Index(i))
				} else if i >= right.Len() {
					out.Index(i).Set(left.Index(i))
				} else {
					out.Index(i).Set(munge(left.Index(i), right.Index(i), f))
				}
			}

			return out
		case reflect.Map:
			var keyType, valueType reflect.Type

			// Check if the keys are compatible
			if right.Type().Key().AssignableTo(left.Type().Key()) {
				keyType = right.Type().Key()
			} else {
				keyType = interfaceType
			}

			// Check if the values are compatible
			if right.Type().Elem().AssignableTo(left.Type().Elem()) {
				valueType = right.Type().Elem()
			} else {
				valueType = interfaceType
			}

			out := reflect.MakeMap(reflect.MapOf(keyType, valueType))

			for _, key := range left.MapKeys() {
				out.SetMapIndex(key, left.MapIndex(key))
			}

			for _, key := range right.MapKeys() {
				l := left.MapIndex(key)
				r := right.MapIndex(key)

				if !l.IsValid() {
					out.SetMapIndex(key, r)
				} else {
					out.SetMapIndex(key, munge(l, r, f))
				}
			}

			return out
		}
	}

	return right
}
