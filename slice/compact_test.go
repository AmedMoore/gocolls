package slice

import (
	"reflect"
	"testing"
)

func TestCompact(t *testing.T) {
	tests := []struct {
		name  string
		input []interface{}
		want  []interface{}
	}{
		{
			name:  "test with nils",
			input: []interface{}{nil, 1, "hello", nil, 3},
			want:  []interface{}{1, "hello", 3},
		},
		{
			name:  "test with zero values",
			input: []interface{}{0, "", false, nil, "golang"},
			want:  []interface{}{"golang"},
		},
		{
			name:  "test with all valid values",
			input: []interface{}{"hello", 1, true},
			want:  []interface{}{"hello", 1, true},
		},
		{
			name:  "test with all invalid values",
			input: []interface{}{nil, "", false, 0},
			want:  []interface{}{},
		},
		{
			name:  "test with empty slice",
			input: []interface{}{},
			want:  []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compact(tt.input)
			if len(tt.want) == 0 {
				if reflect.DeepEqual(got, tt.want) {
					t.Errorf("Compact() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compact() = %v, want %v", got, tt.want)
			}
		})
	}
}
