package maps

import "github.com/amedmoore/gocolls/slice"

// Values returns a slice containing all the values from the given maps.
func Values[Key comparable, Value any](maps ...map[Key]Value) []Value {
	values := make([]Value, 0, slice.Len(maps))
	for _, m := range maps {
		for _, value := range m {
			values = append(values, value)
		}
	}
	return values
}
