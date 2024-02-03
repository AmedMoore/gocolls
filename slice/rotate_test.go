package slice

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		slice    []any
		rotateBy int
		want     []any
	}{
		{
			name:     "rotateByPositiveLessThanLength",
			slice:    []any{1, 2, 3, 4, 5},
			rotateBy: 2,
			want:     []any{3, 4, 5, 1, 2},
		},
		{
			name:     "rotateByNegativeLessThanLength",
			slice:    []any{1, 2, 3, 4, 5},
			rotateBy: -2,
			want:     []any{4, 5, 1, 2, 3},
		},
		{
			name:     "rotateByZero",
			slice:    []any{1, 2, 3, 4, 5},
			rotateBy: 0,
			want:     []any{1, 2, 3, 4, 5},
		},
		{
			name:     "rotateByLength",
			slice:    []any{1, 2, 3, 4, 5},
			rotateBy: 5,
			want:     []any{1, 2, 3, 4, 5},
		},
		{
			name:     "rotateEmptySlice",
			slice:    []any{},
			rotateBy: 3,
			want:     []any{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Rotate(tc.slice, tc.rotateBy)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Rotate(%v, %d) = %v; want %v", tc.slice, tc.rotateBy, got, tc.want)
			}
		})
	}
}
