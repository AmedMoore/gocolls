package slice

import (
	"reflect"
	"testing"
)

func TestBoundedSub(t *testing.T) {
	var tests = []struct {
		name  string
		slice []int
		start int
		end   int
		want  []int
	}{
		{"BothBoundsInbounds", []int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3}},
		{"StartOutOfBounds", []int{1, 2, 3, 4, 5}, -1, 3, []int{1, 2, 3}},
		{"EndOutOfBounds", []int{1, 2, 3, 4, 5}, 1, 6, []int{2, 3, 4, 5}},
		{"BothBoundsOutOfBounds", []int{1, 2, 3, 4, 5}, -1, 6, []int{1, 2, 3, 4, 5}},
		{"BothBoundsSame", []int{1, 2, 3, 4, 5}, 2, 2, []int{}},
		{"EndLowerThanStart", []int{1, 2, 3, 4, 5}, 3, 1, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BoundedSub(tt.slice, tt.start, tt.end)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
