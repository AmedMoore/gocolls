package slice

import (
	"testing"
)

func TestForEach(t *testing.T) {
	cases := []struct {
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
			name:     "One Element",
			input:    []int{5},
			expected: []int{10},
		},
		{
			name:     "Multiple Elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result []int
			ForEach(c.input, func(v int) { result = append(result, v*2) })
			if len(result) != len(c.expected) {
				t.Fatalf("Expected length %d but got %d", len(c.expected), len(result))
			}
			for i, v := range result {
				if v != c.expected[i] {
					t.Fatalf("At index %d: expected %d, got %d", i, c.expected[i], v)
				}
			}
		})
	}
}
