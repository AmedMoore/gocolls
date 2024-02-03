package slice

// Len returns the total length of one or more slices.
func Len[T any](slices ...[]T) int {
	length := 0
	for _, slice := range slices {
		length += len(slice)
	}
	return length
}
