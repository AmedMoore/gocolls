package utils

import (
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want bool
	}{
		{
			name: "All Equal",
			args: []any{1, 1, 1},
			want: true,
		},
		{
			name: "Different Values",
			args: []any{1, 2, 3},
			want: false,
		},
		{
			name: "Single Value",
			args: []any{1},
			want: true,
		},
		{
			name: "Empty Slice",
			args: []any{},
			want: true,
		},
		{
			name: "Equal Slices",
			args: []any{[]int{1, 2, 3}, []int{1, 2, 3}},
			want: true,
		},
		{
			name: "Different Slices",
			args: []any{[]int{1, 2, 3}, []int{1, 2, 4}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args...); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
