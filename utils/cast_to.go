package utils

import "reflect"

// CastTo a generic function that converts a value `from` to type `to`.
func CastTo[T any](from any, to T) T {
	return reflect.ValueOf(from).Convert(reflect.TypeOf(to)).Interface().(T)
}

// CastToType converts a value `from` to the specified reflect
// type `to` and returns the converted value of type `T`.
func CastToType[T any](from any, to reflect.Type) T {
	return reflect.ValueOf(from).Convert(to).Interface().(T)
}
