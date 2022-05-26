package main

import (
	"testing"
)

func Test_Risk_Calculation(t *testing.T) {

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
				input: []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"},
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

func Test_Basin_Calculation(t *testing.T) {

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
				input: []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"},
			},
			result: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hMap := createMap(tt.args.input)
			basinValue := calculateBasinValue(hMap)
			if basinValue != tt.result {
				t.Fatalf("Risk value incorrect!\nExpected: %v\nActual: %v\n", tt.result, basinValue)
			}
		})
	}
}
