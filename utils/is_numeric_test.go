package utils

import (
	"testing"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "Int",
			val:  100,
			want: true,
		},
		{
			name: "Float",
			val:  1.2345,
			want: true,
		},
		{
			name: "Uint",
			val:  uint32(500),
			want: true,
		},
		{
			name: "Complex",
			val:  complex64(10 + 2i),
			want: true,
		},
		{
			name: "Rune",
			val:  'a',
			want: true,
		},
		{
			name: "Byte",
			val:  byte('b'),
			want: true,
		},
		{
			name: "String",
			val:  "not Numeric",
			want: false,
		},
		{
			name: "Tooltip",
			val: struct {
				tip string
			}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.val); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v, type %t", got, tt.want, tt.val)
			}
		})
	}
}
