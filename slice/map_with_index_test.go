package slice

import (
	"reflect"
	"testing"
)

func TestMapWithIndex(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		f      func(int, int) int
		expect []int
	}{
		{
			name:  "IncrementValues",
			input: []int{1, 2, 3, 4, 5},
			f: func(item int, idx int) int {
				return item + idx
			},
			expect: []int{1, 3, 5, 7, 9},
		},
		{
			name:  "MultiplyIndex",
			input: []int{3, 3, 3, 3, 3},
			f: func(item int, idx int) int {
				return item * idx
			},
			expect: []int{0, 3, 6, 9, 12},
		},
		{
			name:  "EmptySlice",
			input: []int{},
			f: func(item int, idx int) int {
				return item * idx
			},
			expect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapWithIndex(tt.input, tt.f)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("MapWithIndex() got = %v, expected = %v", got, tt.expect)
			}
		})
	}
}
