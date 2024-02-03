package slice

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	cases := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "SingleElement",
			slices:   [][]int{{1}},
			expected: []int{1},
		},
		{
			name:     "DuplicateElements",
			slices:   [][]int{{1, 1, 2, 2, 3, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "NoDuplicateElements",
			slices:   [][]int{{1, 2, 3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "MultipleSlices",
			slices:   [][]int{{1, 2}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if output := Unique(c.slices...); !reflect.DeepEqual(output, c.expected) {
				t.Fatalf("Output: %v, Expected: %v", output, c.expected)
			}
		})
	}
}
