package utils

// IsNumeric checks if the given value is a numeric type.
func IsNumeric(val any) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		complex64, complex128:
		return true
	}
	return false
}
