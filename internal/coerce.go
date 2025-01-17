package internal

import (
	"fmt"
	"reflect"
)

type Config struct {
	StripNulls bool
	StringKeys bool
	MapArrays  bool
}

// Coerce wraps its argument,
// guaranteeing that its type and any types in contains
// can safely be formatted by any of the please formatters
func Coerce(in any, config Config) any {
	return newValue(reflect.ValueOf(in), config)
}

var zero = reflect.Value{}

func newValue(in reflect.Value, config Config) any {
	if in == zero {
		if config.StripNulls {
			return "null"
		}
		return nil
	}

	switch in.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
		return in.Interface()
	case reflect.Array, reflect.Slice:
		if config.MapArrays {
			return newMapArray(in, config)
		}
		return newSlice(in, config)
	case reflect.Map:
		if config.StringKeys {
			return newStringKeysMap(in, config)
		}
		return newMap(in, config)
	case reflect.Struct:
		return newStruct(in, config)
	case reflect.Interface, reflect.Pointer:
		return newValue(in.Elem(), config)
	default:
		return fmt.Sprint(in.Interface())
	}
}

func newSlice(in reflect.Value, config Config) []any {
	out := make([]any, in.Len())

	for i := 0; i < in.Len(); i++ {
		out[i] = newValue(in.Index(i), config)
	}

	return out
}

func newMap(in reflect.Value, config Config) map[any]any {
	out := make(map[any]any)

	for key, value := range in.Seq2() {
		realKey := newValue(key, config)
		if config.StringKeys {
			realKey = fmt.Sprint(realKey)
		}
		out[realKey] = newValue(value, config)
	}

	return out
}

func newStringKeysMap(in reflect.Value, config Config) map[string]any {
	out := make(map[string]any)

	keyConfig := config
	keyConfig.StripNulls = true

	for key, value := range in.Seq2() {
		stringKey := fmt.Sprint(newValue(key, keyConfig))

		out[stringKey] = newValue(value, config)
	}

	return out
}

func newStruct(in reflect.Value, config Config) map[string]any {
	out := make(map[string]any)

	t := in.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		out[field.Name] = newValue(in.Field(i), config)
	}

	return out
}

func newMapArray(in reflect.Value, config Config) any {
	out := make(map[int]any)

	for i := 0; i < in.Len(); i++ {
		out[i] = newValue(in.Index(i), config)
	}

	return newValue(reflect.ValueOf(out), config)
}
