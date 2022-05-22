package main

import (
	"testing"
)

func Test_X(t *testing.T) {

	type args struct {
		input []float64
	}
	tests := []struct {
		name   string
		args   args
		result float64
	}{
		{
			name: "Validate test input",
			args: args{
				input: []float64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			},
			result: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalFuel := alignCrabsMedian(tt.args.input)
			if totalFuel != tt.result {
				t.Fatalf("Lanternfish count not correct.\nExpected:\n%v\nActual:\n%v\n",
					tt.result,
					totalFuel)
			}
		})
	}
}
