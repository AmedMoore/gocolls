package slice

import "github.com/amedmoore/gocolls/utils"

// Compact creates a slice with all falsy values removed.
func Compact[T any](slice []T) []T {
	var compacted []T

	for _, element := range slice {
		if !utils.IsFalsy(element) {
			compacted = append(compacted, element)
		}
	}

	return compacted
}
