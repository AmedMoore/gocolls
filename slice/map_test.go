package slice

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	type TestCase struct {
		name     string
		input    []string
		function func(string) string
		expected []string
	}

	var testCases = []TestCase{
		{
			name:     "Map to uppercase",
			input:    []string{"hello", "world"},
			function: strings.ToUpper,
			expected: []string{"HELLO", "WORLD"},
		},
		{
			name:     "Map to empty string",
			input:    []string{"hello", "world"},
			function: func(s string) string { return "" },
			expected: []string{"", ""},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			function: func(s string) string { return s },
			expected: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := Map(testCase.input, testCase.function)
			if !reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("Map(%v, func) = %v, want %v", testCase.input, output, testCase.expected)
			}
		})
	}
}
