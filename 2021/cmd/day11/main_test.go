package main

import (
	"testing"
)

func Test_X(t *testing.T) {

	type args struct {
		input []string
		steps int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "Validate test input",
			args: args{
				input: []string{"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526"},
				steps: 10,
			},
			result: 204,
		},
		{
			name: "Validate test input",
			args: args{
				input: []string{"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526"},
				steps: 100,
			},
			result: 1656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var count = fullDumbo(tt.args.input, tt.args.steps)
			if count != tt.result {
				t.Fatalf("Count incorrect!\nExpected: %d\nActual: %d\n", tt.result, count)
			}
		})
	}
}
