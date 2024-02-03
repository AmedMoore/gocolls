package utils

import "reflect"

// Default returns the default value for a given type.
func Default(value any) any {
	if value == nil || reflect.ValueOf(value).Kind() == reflect.Ptr {
		return nil
	}
	return reflect.Zero(reflect.TypeOf(value)).Interface()
}
