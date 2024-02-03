package slice

import "cmp"

// Sort sorts a slice of elements in ascending order.
func Sort[T cmp.Ordered](slice []T) []T {
	return Order(slice)
}
