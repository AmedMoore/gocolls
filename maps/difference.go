package maps

import "github.com/amedmoore/gocolls/utils"

// Difference returns a new map that contains the key-value pairs from
// map1 that are not present in map2 or have different values.
func Difference[Key comparable, Value any](map1, map2 map[Key]Value) map[Key]Value {
	result := make(map[Key]Value)

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || !utils.Equal(value1, value2) {
			result[key] = value1
		}
	}

	return result
}
