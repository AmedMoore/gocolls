package maps

import "github.com/amedmoore/gocolls/slice"

// Keys returns a slice containing all the keys from the given maps.
func Keys[Key comparable, Value any](maps ...map[Key]Value) []Key {
	keys := make([]Key, 0, slice.Len(maps))
	for _, m := range maps {
		for key := range m {
			keys = append(keys, key)
		}
	}
	return slice.Unique(keys)
}
