package internal_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/please/internal"
)

func TestForceStringKeys(t *testing.T) {
	input := map[any]any{
		13: []any{
			"foo",
			map[int]any{
				0:   "none",
				100: "some",
			},
		},
		66: map[int]any{
			1: []any{"foo", "bar"},
			2: "two",
		},
		nil: "derf",
	}

	expected := map[string]any{
		"13": []any{
			"foo",
			map[string]any{
				"0":   "none",
				"100": "some",
			},
		},
		"66": map[string]any{
			"1": []any{"foo", "bar"},
			"2": "two",
		},
		"null": "derf",
	}

	actual := internal.Coerce(input, internal.Config{StringKeys: true})

	if d := cmp.Diff(expected, actual); d != "" {
		t.Error(d)
	}
}

func TestArraysToMaps(t *testing.T) {
	input := []any{
		[]any{123, 456},
		map[string]any{
			"foo": []any{789},
		},
	}

	expected := map[any]any{
		0: map[any]any{
			0: 123,
			1: 456,
		},
		1: map[any]any{
			"foo": map[any]any{
				0: 789,
			},
		},
	}

	actual := internal.Coerce(input, internal.Config{MapArrays: true})

	if d := cmp.Diff(expected, actual); d != "" {
		t.Error(d)
	}
}
