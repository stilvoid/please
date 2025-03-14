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

func TestNilCoercion(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		config   internal.Config
		expected any
	}{
		{
			name:     "nil with strip nulls",
			input:    nil,
			config:   internal.Config{StripNulls: true},
			expected: "null",
		},
		{
			name:     "nil without strip nulls",
			input:    nil,
			config:   internal.Config{StripNulls: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := internal.Coerce(tt.input, tt.config)
			if d := cmp.Diff(tt.expected, actual); d != "" {
				t.Error(d)
			}
		})
	}
}

func TestStructCoercion(t *testing.T) {
	type testStruct struct {
		Name    string
		Age     int
		IsValid bool
	}

	input := testStruct{
		Name:    "test",
		Age:     25,
		IsValid: true,
	}

	expected := map[string]any{
		"Name":    "test",
		"Age":     25,
		"IsValid": true,
	}

	actual := internal.Coerce(input, internal.Config{})

	if d := cmp.Diff(expected, actual); d != "" {
		t.Error(d)
	}
}

func TestPointerCoercion(t *testing.T) {
	str := "test"
	input := &str

	expected := "test"

	actual := internal.Coerce(input, internal.Config{})

	if d := cmp.Diff(expected, actual); d != "" {
		t.Error(d)
	}
}

func TestDefaultCoercion(t *testing.T) {
	// Test the default case in the switch statement
	// This creates a complex type that doesn't match any of the specific cases
	ch := make(chan int)
	
	actual := internal.Coerce(ch, internal.Config{})
	
	// The result should be a string representation of the channel
	if _, ok := actual.(string); !ok {
		t.Errorf("Expected string output for default case, got %T", actual)
	}
}

func TestMapWithStringKeys(t *testing.T) {
	input := map[any]any{
		13: "value",
		"key": "string key",
	}

	config := internal.Config{StringKeys: true}
	result := internal.Coerce(input, config)

	mapResult, ok := result.(map[string]any)
	if !ok {
		t.Errorf("Expected map[string]any, got %T", result)
	}

	if mapResult["13"] != "value" {
		t.Errorf("Expected value for key '13', got %v", mapResult["13"])
	}
	if mapResult["key"] != "string key" {
		t.Errorf("Expected 'string key' for key 'key', got %v", mapResult["key"])
	}
}
