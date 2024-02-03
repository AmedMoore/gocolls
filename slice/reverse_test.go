package slice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expected    []int
	}{
		{
			description: "Test Normal Array",
			input:       []int{1, 2, 3, 4, 5},
			expected:    []int{5, 4, 3, 2, 1},
		},
		{
			description: "Test Single Element",
			input:       []int{1},
			expected:    []int{1},
		},
		{
			description: "Test No Element",
			input:       []int{},
			expected:    []int{},
		},
		{
			description: "Test Negative Numbers",
			input:       []int{-1, -2, -3},
			expected:    []int{-3, -2, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := Reverse(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("Expected and actual values don't match. Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
