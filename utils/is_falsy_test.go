package utils

import "testing"

func TestIsFalsy(t *testing.T) {
	tests := []struct {
		name   string
		input  any
		expect bool
	}{
		{
			name:   "Test nil value",
			input:  nil,
			expect: true,
		},
		{
			name:   "Test empty string",
			input:  "",
			expect: true,
		},
		{
			name:   "Test zero int",
			input:  0,
			expect: true,
		},
		{
			name:   "Test non-zero int",
			input:  1,
			expect: false,
		},
		{
			name:   "Test non-empty string",
			input:  "Hello",
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsFalsy(test.input); got != test.expect {
				t.Errorf("IsFalsy() = %v, want %v", got, test.expect)
			}
		})
	}
}
