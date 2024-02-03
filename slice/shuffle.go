package slice

import (
	"math/rand"
)

// Shuffle randomly reorders the elements in the given slice. It creates a new
// slice with the same length as the input slice and then shuffles the elements
// by swapping them with a randomly selected element before it. The function
// uses the Fisher-Yates algorithm to ensure a uniform distribution of the
// shuffled elements.
//
// The input slice is not modified by the shuffle operation. Instead, a new
// shuffled slice is created and returned.
//
// Example usage:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	shuffled := Shuffle(numbers)
//	fmt.Println(shuffled) // Example Output: [2, 5, 4, 1, 3]
//
// This function has a time complexity of O(n), where n is the length of the
// input slice.
func Shuffle[T any](slice []T) []T {
	n := len(slice)
	shuffled := make([]T, n)
	copy(shuffled, slice)

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
