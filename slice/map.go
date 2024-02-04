package slice

// Map applies a function to each element of a slice and
// returns a new slice containing the results.
func Map[T any, R any](slice []T, mapFunc func(item T) R) []R {
	if mapFunc == nil {
		return make([]R, 0)
	}
	mapped := make([]R, len(slice))
	for idx, item := range slice {
		mapped[idx] = mapFunc(item)
	}
	return mapped
}
