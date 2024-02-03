package slice

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []interface{}
		item  interface{}
		want  bool
	}{
		{
			name:  "Contains_Int",
			slice: []interface{}{1, 2, 3, 4, 5},
			item:  3,
			want:  true,
		},
		{
			name:  "Does_Not_Contain_Int",
			slice: []interface{}{1, 2, 3, 4, 5},
			item:  6,
			want:  false,
		},
		{
			name:  "Contains_String",
			slice: []interface{}{"a", "b", "c", "d"},
			item:  "c",
			want:  true,
		},
		{
			name:  "Does_Not_Contain_String",
			slice: []interface{}{"a", "b", "c", "d"},
			item:  "e",
			want:  false,
		},
		{
			name:  "Empty_Slice",
			slice: []interface{}{},
			item:  "a",
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.slice, tc.item)

			if got != tc.want {
				t.Errorf("Contains() = %v, want %v", got, tc.want)
			}
		})
	}
}
