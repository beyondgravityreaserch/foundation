package format

import (
	"testing"
)

// Define a custom type that implements the Stringer interface for testing.
type customStringer struct {
	value string
}

func (cs customStringer) String() string {
	return cs.value
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "nil value",
			input:    []any{nil},
			expected: "nil",
		},
		{
			name:     "string value",
			input:    []any{"hello", " ", "world"},
			expected: "hello world",
		},
		{
			name:     "boolean values",
			input:    []any{true, false},
			expected: "truefalse",
		},
		{
			name:     "integer values",
			input:    []any{42, int8(-8), int16(32000)},
			expected: "42-832000",
		},
		{
			name:     "unsigned integer values",
			input:    []any{uint(42), uint16(32000)},
			expected: "4232000",
		},
		{
			name:     "floating-point values",
			input:    []any{3.14, float32(2.71)},
			expected: "3.1400002.710000",
		},
		{
			name:     "complex numbers",
			input:    []any{complex(1, 2), complex128(complex(3.1, 4.2))},
			expected: "(1+2i)(3.1+4.2i)",
		},
		{
			name:     "error type",
			input:    []any{testError("this is an error")},
			expected: "this is an error",
		},
		{
			name:     "Stringer type",
			input:    []any{customStringer{"custom value"}},
			expected: "custom value",
		},
		{
			name:     "default type",
			input:    []any{[]int{1, 2, 3}},
			expected: "[1 2 3]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := ToString(test.input...)
			if output != test.expected {
				t.Errorf("ToString(%v) = %q; want %q", test.input, output, test.expected)
			}
		})
	}
}

// testError is a custom error type for testing.
type testError string

func (e testError) Error() string {
	return string(e)
}
