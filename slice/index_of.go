package slice

import "github.com/amedmoore/gocolls/utils"

// NotFound represents the value -1 which is commonly used to indicate
// that a requested item or index is not found in a collection or container.
const NotFound = -1

// IndexOf finds the index of an item in a slice.
// If fromIndex is provided, the search starts from that index.
// Returns -1 if the item is not found.
func IndexOf[T any](slice []T, item T, fromIndex ...int) int {
	startIndex := fromIndexOrDefault(fromIndex...)

	if startIndex < 0 || startIndex >= len(slice) {
		return NotFound
	}

	for i := startIndex; i < len(slice); i++ {
		if utils.Equal(slice[i], item) {
			return i
		}
	}

	return NotFound
}

func fromIndexOrDefault(index ...int) int {
	if len(index) > 0 {
		return index[len(index)-1]
	}
	return 0
}
