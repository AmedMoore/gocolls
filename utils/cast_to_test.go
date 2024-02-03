package utils

import (
	"reflect"
	"testing"
)

func TestCastTo(t *testing.T) {
	tests := []struct {
		name    string
		from    interface{}
		to      interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "To int from int",
			from:    123,
			to:      0,
			want:    123,
			wantErr: false,
		},
		{
			name:    "To string from string",
			from:    "hello",
			to:      "",
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "To float64 from float64",
			from:    1.23,
			to:      0.0,
			want:    1.23,
			wantErr: false,
		},
		{
			name:    "To bool from bool",
			from:    true,
			to:      false,
			want:    true,
			wantErr: false,
		},
		{
			name:    "To int from incompatible type",
			from:    "hello",
			to:      0,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tt.wantErr {
					t.Errorf("unexpected panic occurred")
				}
			}()

			got := CastTo(tt.from, tt.to)
			if tt.want != got {
				t.Errorf("Expected '%v', but got '%v'", tt.want, got)
			}
		})
	}
}

func TestCastToType(t *testing.T) {
	type testCastToTypeStruct struct {
		IntegerValue  int
		RandomFloat32 float32
	}

	cases := []struct {
		Name          string
		From          any
		To            reflect.Type
		Expected      any
		ExpectedError bool
	}{
		{
			Name:     "Cast int to int success",
			From:     456,
			To:       reflect.TypeOf(1),
			Expected: 456,
		},
		{
			Name:          "Cast string to int error",
			From:          "test string",
			To:            reflect.TypeOf(1),
			ExpectedError: true,
		},
		{
			Name:     "Cast int to float32 success",
			From:     123,
			To:       reflect.TypeOf(float32(1.23)),
			Expected: float32(123),
		},
		{
			Name:     "Cast string to string",
			From:     "test string",
			To:       reflect.TypeOf(""),
			Expected: "test string",
		},
		{
			Name:          "Cast string to custom struct error",
			From:          "test string",
			To:            reflect.TypeOf(testCastToTypeStruct{}),
			Expected:      testCastToTypeStruct{},
			ExpectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.ExpectedError {
					t.Errorf("The code panicked: \n%s\n when it should not have", r)
				} else if r == nil && c.ExpectedError {
					t.Errorf("The code did not panic when it should have")
				}
			}()

			res := CastToType[any](c.From, c.To)

			if c.ExpectedError {
				return
			}

			if !reflect.DeepEqual(res, c.Expected) {
				t.Errorf(`Got: %v, expected: %v`, res, c.Expected)
			}
		})
	}
}
