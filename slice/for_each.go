package slice

// ForEach applies a given function to each element of a slice.
func ForEach[T any](slice []T, iterator func(item T)) {
	if iterator == nil {
		return
	}
	for _, item := range slice {
		iterator(item)
	}
}
