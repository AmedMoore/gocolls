package slice

import (
	"reflect"
	"testing"
)

func TestSub(t *testing.T) {
	type args struct {
		slice []any
		start int
		end   int
	}

	tests := []struct {
		name          string
		args          args
		expectedSlice []any
	}{
		{
			name: "Test with integers",
			args: args{
				slice: []any{1, 2, 3, 4, 5, 6},
				start: 2,
				end:   4,
			},
			expectedSlice: []any{3, 4},
		},
		{
			name: "Test with strings",
			args: args{
				slice: []any{"one", "two", "three", "four", "five"},
				start: 1,
				end:   3,
			},
			expectedSlice: []any{"two", "three"},
		},
		{
			name: "Test with single element",
			args: args{
				slice: []any{1},
				start: 0,
				end:   1,
			},
			expectedSlice: []any{1},
		},
		{
			name: "Test with empty slice",
			args: args{
				slice: []any{},
				start: 0,
				end:   0,
			},
			expectedSlice: []any{},
		},
		{
			name: "Test with start equals end",
			args: args{
				slice: []any{1, 2, 3},
				start: 2,
				end:   2,
			},
			expectedSlice: []any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.slice, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.expectedSlice) {
				t.Errorf("Sub() = %v, expected %v", got, tt.expectedSlice)
			}
		})
	}
}
