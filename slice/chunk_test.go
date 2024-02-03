package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		chunkSize int
		expected  [][]int
	}{
		{
			name:      "Normal case",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: 2,
			expected:  [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:      "Empty slice",
			input:     []int{},
			chunkSize: 3,
			expected:  [][]int{},
		},
		{
			name:      "Chunk size larger than slice",
			input:     []int{1, 2, 3},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3}},
		},
		{
			name:      "Chunk size equal to slice size",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name:      "Chunk size one",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: 1,
			expected:  [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.input, tt.chunkSize)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Chunk() = %v, want %v", got, tt.expected)
			}
		})
	}
}
