package slice

// ForEach applies a given function to each element of a slice.
func ForEach[T any](slice []T, mapFunc func(item T)) {
	for _, item := range slice {
		mapFunc(item)
	}
}
