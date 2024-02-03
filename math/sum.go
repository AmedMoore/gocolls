package math

import "github.com/amedmoore/gocolls/utils"

// Sum calculates the sum of elements in one or more slices
func Sum[T Numeric](slices ...[]T) T {
	zero := utils.CastToType[T](0, utils.ElemType(slices))

	if len(slices) == 0 || len(slices[0]) == 0 {
		return zero
	}

	sum := zero

	for _, slice := range slices {
		for _, value := range slice {
			sum += value
		}
	}

	return sum
}
