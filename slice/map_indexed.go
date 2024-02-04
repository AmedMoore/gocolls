package slice

// MapIndexed applies the given map function to each item in the slice along with its index.
func MapIndexed[T any, R any](slice []T, mapFunc func(item T, index int) R) []R {
	mapped := make([]R, len(slice))
	for idx, item := range slice {
		mapped[idx] = mapFunc(item, idx)
	}
	return mapped
}
