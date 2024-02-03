package slice

import (
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:  "All elements pass",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool {
				return n > 0
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "No elements pass",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool {
				return n > 5
			},
			want: []int{},
		},
		{
			name:  "Some elements pass",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool {
				return n%2 == 0
			},
			want: []int{2, 4},
		},
		{
			name:  "Empty slice",
			slice: []int{},
			predicate: func(n int) bool {
				return n > 0
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.slice, tt.predicate); !isEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, item := range slice1 {
		if item != slice2[i] {
			return false
		}
	}
	return true
}
