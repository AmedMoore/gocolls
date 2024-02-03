package maps

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		map1 map[int]string
		map2 map[int]string
		want map[int]string
	}{
		{
			name: "Empty maps",
			map1: map[int]string{},
			map2: map[int]string{},
			want: map[int]string{},
		},
		{
			name: "Non-intersecting maps",
			map1: map[int]string{1: "a", 2: "b"},
			map2: map[int]string{3: "c", 4: "d"},
			want: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "Intersecting maps different values",
			map1: map[int]string{1: "a", 2: "b"},
			map2: map[int]string{1: "b", 2: "a"},
			want: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "Intersecting maps similar values",
			map1: map[int]string{1: "a", 2: "b"},
			map2: map[int]string{1: "a", 3: "c"},
			want: map[int]string{2: "b"},
		},
		{
			name: "Map1 is subset of Map2",
			map1: map[int]string{1: "a", 2: "b"},
			map2: map[int]string{1: "a", 2: "b", 3: "c"},
			want: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.map1, tt.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
