package slice

// Merge takes in multiple slices of any type and merges them into a single slice.
//
// Example usage:
//
//	slice1 := []int{1, 2, 3}
//	slice2 := []int{4, 5, 6}
//	merged := Merge(slice1, slice2) // Result: [1, 2, 3, 4, 5, 6]
//
// Parameters:
//
//	slices: Multiple slices of any type to be merged.
//
// Returns:
//
//	The merged slice containing elements from all the input slices.
func Merge[T any](slices ...[]T) []T {
	merged := make([]T, Len(slices...))
	index := 0
	for _, slice := range slices {
		copy(merged[index:], slice)
		index += len(slice)
	}
	return merged
}
