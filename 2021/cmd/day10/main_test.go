package main

import (
	"testing"
)

func Test_Part_1(t *testing.T) {

	type args struct {
		input []string
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "Validate test input",
			args: args{
				input: []string{"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]"},
			},
			result: 26397,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := getErrorScore(tt.args.input)
			if score != tt.result {
				t.Fatalf("Score incorrect!\nExpected: %d\nActual: %d\n", tt.result, score)
			}
		})
	}
}
