package slice

import (
	"reflect"
	"testing"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		orders []SortingOrder
		want   []int
	}{
		{
			name:   "Ascending",
			slice:  []int{4, 3, 2, 1},
			orders: []SortingOrder{Asc},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "Descending",
			slice:  []int{1, 2, 3, 4},
			orders: []SortingOrder{Desc},
			want:   []int{4, 3, 2, 1},
		},
		{
			name:   "Without Orders",
			slice:  []int{3, 2, 4, 1},
			orders: []SortingOrder{},
			want:   []int{1, 2, 3, 4}, // Assuming the default order is Ascending
		},
		{
			name:   "With Duplicate Values",
			slice:  []int{5, 2, 2, 5},
			orders: []SortingOrder{Asc},
			want:   []int{2, 2, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Order(tt.slice, tt.orders...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
