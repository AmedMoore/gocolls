package slice

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		item  int
		index int
		want  int
	}{
		{
			name:  "Positive_Index_Within_Range",
			slice: []int{1, 2, 3, 4, 5},
			item:  3,
			index: 0,
			want:  2,
		},
		{
			name:  "Positive_Index_Outside_Range",
			slice: []int{1, 2, 3, 4, 5},
			item:  5,
			index: 5,
			want:  NotFound,
		},
		{
			name:  "Negative_Index",
			slice: []int{1, 2, 3, 4, 5},
			item:  2,
			index: -1,
			want:  NotFound,
		},
		{
			name:  "No_Index_Price",
			slice: []int{1, 2, 3, 4, 5},
			item:  2,
			want:  1,
		},
		{
			name:  "Empty_Slice",
			slice: []int{},
			item:  2,
			index: 0,
			want:  NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.slice, tt.item, tt.index); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
