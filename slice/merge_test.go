package slice

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type testCase struct {
		name     string
		input    [][]int
		expected []int
	}

	tests := []testCase{
		{
			name:     "empty slices",
			input:    [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "single element slices",
			input:    [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "multi element slices",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "multi element slices with duplicates",
			input:    [][]int{{1, 2, 2}, {2, 3, 3}, {3, 4, 4}},
			expected: []int{1, 2, 2, 2, 3, 3, 3, 4, 4},
		},
		{
			name:     "empty slice followed by non empty",
			input:    [][]int{{}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "non empty slice followed by empty",
			input:    [][]int{{1, 2, 3}, {}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty input",
			input:    [][]int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Merge(test.input...)
			if !reflect.DeepEqual(got, test.expected) {
				t.Fatalf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
