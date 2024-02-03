package utils

import (
	"reflect"
	"testing"
)

func TestElemType(t *testing.T) {
	tests := []struct {
		name string
		fn   func() interface{}
		want reflect.Type
	}{
		{
			name: "IntegerSlice",
			fn: func() interface{} {
				return ElemType([]int{1, 2, 3})
			},
			want: reflect.TypeOf(0),
		},
		{
			name: "StringSlice",
			fn: func() interface{} {
				return ElemType([]string{"a", "b", "c"})
			},
			want: reflect.TypeOf(""),
		},
		{
			name: "FloatSlice",
			fn: func() interface{} {
				return ElemType([]float64{1.0, 2.0, 3.0})
			},
			want: reflect.TypeOf(0.0),
		},
		{
			name: "StructSlice",
			fn: func() interface{} {
				return ElemType([]struct{ A, B string }{{"A", "B"}})
			},
			want: reflect.TypeOf(struct{ A, B string }{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}
