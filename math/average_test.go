package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type test[T Numeric] struct {
		name     string
		input    [][]T
		expected T
	}

	intTests := []test[int]{
		{
			name:     "no_slices",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty_slice",
			input:    [][]int{{}},
			expected: 0,
		},
		{
			name:     "positive_numbers",
			input:    [][]int{{1, 2, 3, 4, 5}},
			expected: 3,
		},
		{
			name:     "multi_dimensional",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 3,
		},
		{
			name:     "negative_numbers",
			input:    [][]int{{-1, -2, -3, -4, -5}},
			expected: -3,
		},
	}

	t.Run("int", func(t *testing.T) {
		for _, tc := range intTests {
			t.Run(tc.name, func(t *testing.T) {
				result := Average(tc.input...)
				if result != tc.expected {
					t.Errorf("Average of %v = %d; want %d", tc.input, result, tc.expected)
				}
			})
		}
	})

	floatTests := []test[float64]{
		{
			name:     "no_slices",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty_slice",
			input:    [][]float64{{}},
			expected: 0,
		},
		{
			name:     "positive_numbers",
			input:    [][]float64{{1.0, 2.0, 3.0, 4.0, 5.0}},
			expected: 3.0,
		},
		{
			name:     "multi_dimensional",
			input:    [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}},
			expected: 3.5,
		},
		{
			name:     "negative_numbers",
			input:    [][]float64{{-1.0, -2.0, -3.0, -4.0, -5.0}},
			expected: -3.0,
		},
	}

	t.Run("float64", func(t *testing.T) {
		for _, tc := range floatTests {
			t.Run(tc.name, func(t *testing.T) {
				if result := Average(tc.input...); result != tc.expected {
					t.Errorf("Average of %v = %v; want %v", tc.input, result, tc.expected)
				}
			})
		}
	})
}
