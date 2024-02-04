package slice

// Map applies a function to each element of a slice and
// returns a new slice containing the results.
func Map[T any, R any](slice []T, mapFunc func(item T) R) []R {
	mapped := make([]R, len(slice))
	for idx, item := range slice {
		mapped[idx] = mapFunc(item)
	}
	return mapped
}
