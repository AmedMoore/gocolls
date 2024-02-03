package slice

import "testing"

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{
			name: "Single non-empty slice",
			args: [][]int{{1, 2, 3}},
			want: 3,
		},
		{
			name: "Multiple non-empty slices",
			args: [][]int{{1, 2}, {3, 4, 5, 6}, {7}},
			want: 7,
		},
		{
			name: "Single empty slice",
			args: [][]int{{}},
			want: 0,
		},
		{
			name: "Multiple empty slices",
			args: [][]int{{}, {}, {}},
			want: 0,
		},
		{
			name: "No slice",
			args: [][]int{},
			want: 0,
		},
		{
			name: "Empty and non-empty slices",
			args: [][]int{{}, {1, 2}, {}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.args...); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
