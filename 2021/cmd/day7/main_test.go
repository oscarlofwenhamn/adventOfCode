package main

import (
	"testing"
)

func Test_X(t *testing.T) {

	type args struct {
		input []int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "Validate test input",
			args: args{
				input: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			},
			result: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
