package slice

import (
	"testing"
)

func TestRegulateBounds(t *testing.T) {
	var tests = []struct {
		name     string
		start    int
		end      int
		l        int
		expected [2]int
	}{
		{name: "valid bounds", start: 3, end: 5, l: 8, expected: [2]int{3, 5}},
		{name: "negative start bounds", start: -3, end: 5, l: 8, expected: [2]int{0, 5}},
		{name: "end bounds greater than length", start: 3, end: 12, l: 8, expected: [2]int{3, 8}},
		{name: "end less than start", start: 5, end: 3, l: 8, expected: [2]int{5, 5}},
		{name: "zero length", start: 5, end: 3, l: 0, expected: [2]int{0, 0}},
		{name: "zero bounds and length", start: 0, end: 0, l: 0, expected: [2]int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd := RegulateBounds(tt.start, tt.end, tt.l)
			if gotStart != tt.expected[0] || gotEnd != tt.expected[1] {
				t.Errorf("RegulateBounds(%d, %d, %d) = %d, %d; want %d, %d",
					tt.start, tt.end, tt.l, gotStart, gotEnd, tt.expected[0], tt.expected[1])
			}
		})
	}
}
