package slice

import (
	"testing"
)

func TestForForEachWithStop(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	want := 3
	got := 0
	ForEachWithStop(slice, func(item int, stop func()) {
		got = item
		if item == want {
			stop()
		}
	})
	if got != want {
		t.Errorf("ForEachWithStop() got %v, want %v", got, want)
	}
}
