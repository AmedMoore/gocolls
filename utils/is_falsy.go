package utils

// IsFalsy checks if a value is falsy.
func IsFalsy(value any) bool {
	return Equal(value, Default(value))
}
