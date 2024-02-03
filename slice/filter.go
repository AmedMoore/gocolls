package slice

// Filter takes a slice and a predicate function as input parameters
// and returns a new slice containing only the elements that satisfy
// the predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
