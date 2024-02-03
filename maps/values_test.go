package maps

import (
	"github.com/amedmoore/gocolls/slice"
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	type inputData struct {
		maps []map[int]string
		want []string
	}
	tests := []struct {
		name string
		data inputData
	}{
		{
			"Empty map",
			inputData{
				maps: []map[int]string{{}},
				want: []string{},
			},
		},
		{
			"Single map",
			inputData{
				maps: []map[int]string{{1: "a", 2: "b"}},
				want: []string{"a", "b"},
			},
		},
		{
			"Multiple map",
			inputData{
				maps: []map[int]string{{1: "a", 2: "b"}, {3: "c"}},
				want: []string{"a", "b", "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.Sort(Values(tt.data.maps...))
			if !reflect.DeepEqual(got, tt.data.want) {
				t.Errorf("Values() = %v, want %v", got, tt.data.want)
			}
		})
	}
}
