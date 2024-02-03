package slice

// Sub takes a slice and returns a sub-slice from the specified start index to the end index (exclusive).
// The start index is inclusive, while the end index is exclusive.
//
// Example usage:
//
//	slice := []int{1, 2, 3, 4, 5, 6}
//	subSlice := Sub(slice, 2, 4) // subSlice = [3, 4]
func Sub[T any](slice []T, start, end int) []T {
	return slice[start:end]
}
