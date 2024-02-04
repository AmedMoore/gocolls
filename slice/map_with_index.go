package slice

// MapWithIndex applies the given map function to each item in the slice along with its index.
func MapWithIndex[T any, R any](slice []T, mapFunc func(item T, index int) R) []R {
	if mapFunc == nil {
		return make([]R, 0)
	}
	mapped := make([]R, len(slice))
	for idx, item := range slice {
		mapped[idx] = mapFunc(item, idx)
	}
	return mapped
}
