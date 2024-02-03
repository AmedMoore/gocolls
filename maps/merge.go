package maps

// Merge takes multiple maps as input and merges them into a single map.
func Merge[Key comparable, Value any](maps ...map[Key]Value) map[Key]Value {
	merged := make(map[Key]Value)
	for _, m := range maps {
		for key, value := range m {
			merged[key] = value
		}
	}
	return merged
}
