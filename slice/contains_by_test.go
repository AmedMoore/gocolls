package slice

import "testing"

func TestContainsBy(t *testing.T) {
	cases := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "EmptySlice",
			slice:     []int{},
			predicate: func(i int) bool { return i == 5 },
			want:      false,
		},
		{
			name:      "NotContains",
			slice:     []int{1, 2, 3, 4, 6},
			predicate: func(i int) bool { return i == 5 },
			want:      false,
		},
		{
			name:      "Contains",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i == 5 },
			want:      true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ContainsBy(tc.slice, tc.predicate); got != tc.want {
				t.Errorf("ContainsBy() = %v, want %v", got, tc.want)
			}
		})
	}
}
