package utils

import "reflect"

// ElemType returns the type of elements in a given slice or pointer type.
func ElemType[T any](slice []T) reflect.Type {
	for t := reflect.TypeOf(slice); ; {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice:
			t = t.Elem()
		default:
			return t
		}
	}
}
