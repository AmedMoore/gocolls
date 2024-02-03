package slice

// Unique takes in one or more slices of comparable elements and returns a new slice
// containing only the unique elements from the input slices.
func Unique[T comparable](slices ...[]T) []T {
	uniqueElements := make(map[T]bool)
	var result []T

	for _, slice := range slices {
		for _, element := range slice {
			if _, exists := uniqueElements[element]; !exists {
				uniqueElements[element] = true
				result = append(result, element)
			}
		}
	}

	return result
}
