package slice

// Rotate rotates the elements of a slice by a given number of positions.
// The rotate operation wraps around the slice, so elements shifted beyond
// the end of the slice are moved to the beginning.
// The original slice is not modified and a new rotated slice is returned.
//
// Parameters:
//   - slice: The slice to rotate.
//   - i: The number of positions to rotate the slice by. Positive values
//     rotate the slice to the right, while negative values rotate it
//     to the left.
//
// Returns:
//
//	The rotated slice.
func Rotate[T any](slice []T, i int) []T {
	length := len(slice)
	if length == 0 {
		return slice
	}
	i = (i%length + length) % length
	return append(slice[i:], slice[:i]...)
}
