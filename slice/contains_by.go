package slice

// ContainsBy checks if an item is present in a given slice.
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	if predicate == nil {
		return false
	}
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}
