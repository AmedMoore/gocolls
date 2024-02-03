package maps

import "github.com/amedmoore/gocolls/utils"

// Intersection returns a new map that contains the key-value pairs that are common between map1 and map2.
//
// Example Usage:
//
//	map1 := map[int]int{1: 1, 2: 2, 3: 3}
//	map2 := map[int]int{1: 1, 2: 2, 3: 4}
//	result := Intersection(map1, map2) // map[int]int{1: 1, 2: 2}
func Intersection[Key comparable, Value any](map1, map2 map[Key]Value) map[Key]Value {
	result := make(map[Key]Value)

	for key, value1 := range map1 {
		if value2, ok := map2[key]; ok && utils.Equal(value1, value2) {
			result[key] = value1
		}
	}

	return result
}
