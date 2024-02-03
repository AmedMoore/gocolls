package utils

import (
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want any
	}{
		{"String", "Hello", ""},
		{"Integer", 10, 0},
		{"Float", 1.5, 0.0},
		{"Boolean", true, false},
		{"Empty String", "", ""},
		{"Zero Integer", 0, 0},
		{"Zero Float", 0.0, 0.0},
		{"False Boolean", false, false},
		{"Nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
