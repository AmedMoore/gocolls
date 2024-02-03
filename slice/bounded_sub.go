package slice

// BoundedSub returns a sub-slice of a given slice, bounded by start and end indices.
// The start index is inclusive, while the end index is exclusive.
// If the start index is negative, it will be set to 0.
// If the end index is greater than the length of the slice, it will be set to the length of the slice.
// If the end index is less than the start index, the sub-slice will be empty.
//
// Example usage:
//
//	slice := []int{1, 2, 3, 4, 5}
//	start := 1
//	end := 3
//	result := BoundedSub(slice, start, end) // []int{2, 3}
func BoundedSub[T any](slice []T, start, end int) []T {
	length := len(slice)
	start, end = RegulateBounds(start, end, length)
	return slice[start:end]
}
