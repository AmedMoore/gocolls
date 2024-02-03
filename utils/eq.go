package utils

import "reflect"

// Equal runs a deep equality check on all items in the variadic argument list.
// This function is designed to compare complex values, for comparable
// types please use the comparison operators == and !=.
func Equal(args ...any) bool {
	if len(args) < 2 {
		return true
	}
	first := args[0]
	for _, arg := range args[1:] {
		if !reflect.DeepEqual(first, arg) {
			return false
		}
	}
	return true
}
