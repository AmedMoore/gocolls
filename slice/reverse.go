package slice

// Reverse reverses the order of elements in a slice.
func Reverse[T any](slice []T) []T {
	n := len(slice)
	reversed := make([]T, n)

	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		reversed[i] = slice[j]
	}

	return reversed
}
