package maps

import (
	"reflect"
	"testing"

	"github.com/amedmoore/gocolls/slice"
)

func TestKeys(t *testing.T) {
	type Test struct {
		name   string
		input  []map[int]string
		output []int
	}
	tests := []Test{
		{
			name:   "SingleMapWithSingleEntry",
			input:  []map[int]string{{1: "one"}},
			output: []int{1},
		},
		{
			name:   "MultipleMaps",
			input:  []map[int]string{{1: "one"}, {2: "two"}},
			output: []int{1, 2},
		},
		{
			name:   "MultipleEntriesInSingleMap",
			input:  []map[int]string{{1: "one", 2: "two"}},
			output: []int{1, 2},
		},
		{
			name:   "DuplicateKeysInMultipleMaps",
			input:  []map[int]string{{1: "one"}, {1: "uno"}},
			output: []int{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slice.Sort(Keys(test.input...))
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %v, but got %v", test.output, result)
			}
		})
	}
}
