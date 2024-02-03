package slice

import (
	"cmp"
	"sort"
)

// SortingOrder represents the order in which to sort a collection of items.
//
// SortingOrder values can be "asc" (ascending) or "desc" (descending).
type SortingOrder string

const (
	Asc  SortingOrder = "asc"
	Desc SortingOrder = "desc"
)

// Order sorts a given slice of elements in ascending or descending order based on the provided sorting order.
//
// The slice is not modified, and a new sorted slice is returned.
//
// The sorting order can be specified as optional arguments.
//
// If no sorting order is provided, the default sorting order is ascending.
func Order[T cmp.Ordered](slice []T, orders ...SortingOrder) []T {
	order := orderOrDefault(orders...)
	sorted := make([]T, len(slice))
	copy(sorted, slice)
	switch order {
	case Asc:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
	case Desc:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		})
	}
	return sorted
}

func orderOrDefault(order ...SortingOrder) SortingOrder {
	if len(order) > 0 {
		return order[len(order)-1]
	}
	return Asc
}
