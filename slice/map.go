package slice

// Map applies a function to each element of a slice and
// returns a new slice containing the results.
func Map[T any, R any](slice []T, mapFunc func(T) R) []R {
	mapped := make([]R, 0)
	for _, item := range slice {
		mapped = append(mapped, mapFunc(item))
	}
	return mapped
}
