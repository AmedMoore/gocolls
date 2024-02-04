package slice

// ForEachIndexed iterates over a slice and applies the specified map function
// to each element, providing the element value and its index as arguments
// to the map function.
func ForEachIndexed[T any](slice []T, mapFunc func(item T, index int)) {
	for idx, item := range slice {
		mapFunc(item, idx)
	}
}
