package slice

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	type test struct {
		name   string
		slices [][]int
		want   []int
	}

	tests := []test{
		{
			name:   "Empty Slices",
			slices: [][]int{},
			want:   []int{},
		},
		{
			name:   "Single Element Slices",
			slices: [][]int{{1}, {2}, {3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "Multi Element Slices",
			slices: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Mixed Size Slices",
			slices: [][]int{{1}, {2, 3}, {4, 5, 6}},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Slices Containing Zero",
			slices: [][]int{{1, 2, 0}, {0, 3, 4}, {5, 6, 0}},
			want:   []int{1, 2, 0, 0, 3, 4, 5, 6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flatten(tt.slices)
			if len(tt.want) == 0 {
				if reflect.DeepEqual(got, tt.want) {
					t.Errorf("Flatten() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
