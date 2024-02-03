package slice

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	// set up tests cases
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty Slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single Element",
			input:    []int{20},
			expected: []int{20},
		},
		{
			name:     "Multiple Elements Ascending",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Multiple Elements Descending",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Multiple Comparable Elements",
			input:    []int{5, 3, 4, 5, 2},
			expected: []int{2, 3, 4, 5, 5},
		},
	}

	// perform tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("for case [%s], expected: %v, but got: %v", tc.name, tc.expected, result)
			}
		})
	}
}
