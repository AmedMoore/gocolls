package slice

import "github.com/amedmoore/gocolls/utils"

// Contains checks if an item is present in a given slice.
func Contains[T any](slice []T, item T) bool {
	for _, v := range slice {
		if utils.Equal(v, item) {
			return true
		}
	}
	return false
}
