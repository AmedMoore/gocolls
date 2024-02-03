package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   int
	}{
		{
			name:   "No slices",
			slices: [][]int{},
			want:   0,
		},
		{
			name:   "Empty slice",
			slices: [][]int{{}},
			want:   0,
		},
		{
			name:   "Single slice",
			slices: [][]int{{1, 2, 3, 4}},
			want:   10,
		},
		{
			name:   "Multiple slices",
			slices: [][]int{{1, 2}, {3, 4}},
			want:   10,
		},
		{
			name:   "Multiple slices, different lengths",
			slices: [][]int{{1, 2}, {3, 4, 5}},
			want:   15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Sum(tt.slices...); result != tt.want {
				t.Errorf("Sum of %v = %v; want %v", tt.slices, result, tt.want)
			}
		})
	}
}
