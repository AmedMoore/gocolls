package slice

// ForEachWithIndex iterates over a slice and applies the specified map function
// to each element, providing the element value and its index as arguments
// to the map function.
func ForEachWithIndex[T any](slice []T, iterator func(item T, index int)) {
	if iterator == nil {
		return
	}
	for idx, item := range slice {
		iterator(item, idx)
	}
}
