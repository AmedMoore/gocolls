package maps

import (
	"github.com/amedmoore/gocolls/utils"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		map1     map[int]int
		map2     map[int]int
		expected map[int]int
	}{
		{
			name:     "Empty Maps",
			map1:     map[int]int{},
			map2:     map[int]int{},
			expected: map[int]int{},
		},
		{
			name:     "Non Empty Map and Empty Map",
			map1:     map[int]int{1: 1, 2: 2},
			map2:     map[int]int{},
			expected: map[int]int{},
		},
		{
			name:     "No Common Elements",
			map1:     map[int]int{1: 1, 2: 2},
			map2:     map[int]int{3: 3, 4: 4},
			expected: map[int]int{},
		},
		{
			name:     "Common Elements But Different Values",
			map1:     map[int]int{1: 1, 2: 2},
			map2:     map[int]int{1: 2, 2: 3},
			expected: map[int]int{},
		},
		{
			name:     "Common Elements Same Values",
			map1:     map[int]int{1: 1, 2: 2, 3: 3},
			map2:     map[int]int{1: 1, 2: 2, 3: 4},
			expected: map[int]int{1: 1, 2: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.map1, tt.map2)
			if !utils.Equal(got, tt.expected) {
				t.Errorf("Intersection got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
