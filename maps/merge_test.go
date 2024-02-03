package maps

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		maps   []map[int]int
		result map[int]int
	}{
		{
			name:   "No maps",
			maps:   []map[int]int{},
			result: map[int]int{},
		},
		{
			name: "Single map",
			maps: []map[int]int{{1: 1}},
			result: map[int]int{1: 1},
		},
		{
			name: "Two equal maps",
			maps: []map[int]int{{1: 1, 2: 2}, {1: 1, 2: 2}},
			result: map[int]int{1: 1, 2: 2},
		},
		{
			name: "Two different maps with different keys",
			maps: []map[int]int{{1: 1}, {2: 2}},
			result: map[int]int{1: 1, 2: 2},
		},
		{
			name: "Two different maps with same keys",
			maps: []map[int]int{{1: 1}, {1: 2}},
			result: map[int]int{1: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.maps...); !reflect.DeepEqual(got, tt.result) {
				t.Errorf("Merge() = %v, want %v", got, tt.result)
			}
		})
	}
}
