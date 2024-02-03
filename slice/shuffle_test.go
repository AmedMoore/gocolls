package slice

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		in   []int
	}{
		{
			name: "normal",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "empty",
			in:   []int{},
		},
		{
			name: "one_element",
			in:   []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shuffle(tt.in)
			if len(tt.in) < 2 {
				if !reflect.DeepEqual(got, tt.in) {
					t.Errorf("Shuffle() = %v, want non shuffled slice %v", got, tt.in)
				}
			}
			if len(tt.in) >= 2 && reflect.DeepEqual(got, tt.in) {
				t.Errorf("Shuffle() = %v, want shuffled slice %v", got, tt.in)
			}
			if len(got) != len(tt.in) {
				t.Errorf("Shuffle() length = %v, want %v", len(got), len(tt.in))
			}
			m := make(map[int]int)
			for _, v := range tt.in {
				m[v]++
			}
			for _, v := range got {
				m[v]--
			}
			for k, v := range m {
				if v != 0 {
					t.Errorf("Shuffle() contents error = element %v: pre-shuffle count %v, post-shuffle count %v", k, v, 0)
				}
			}
		})
	}
}
