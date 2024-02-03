package slice

// Flatten takes a slice of slices of any type and returns a flattened slice of the same type.
func Flatten[T any](slices [][]T) []T {
	var result []T

	for _, slice := range slices {
		result = append(result, slice...)
	}

	return result
}
