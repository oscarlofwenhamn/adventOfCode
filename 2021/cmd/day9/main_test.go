package main

import (
	"testing"
)

func Test_X(t *testing.T) {

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
				input: []string{},
			},
			result: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hMap := createMap(tt.args.input)
			risk := calculateRisk(hMap)
			if risk != tt.result {
				t.Fatalf("Risk value incorrect!\nExpected: %v\nActual: %v\n", tt.result, risk)
			}
		})
	}
}
