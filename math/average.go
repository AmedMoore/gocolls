package math

import (
	"github.com/amedmoore/gocolls/slice"
	"github.com/amedmoore/gocolls/utils"
)

// Average compute the average value of the elements in the given slices.
func Average[T Numeric](slices ...[]T) T {
	sum := Sum(slices...)
	if sum == 0 {
		return 0
	}
	total := utils.CastTo(slice.Len(slices...), sum)
	return sum / total
}
