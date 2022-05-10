package main

import (
	"testing"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/day6/model"
)

func Test_spendDays(t *testing.T) {

	type args struct {
		initialValues []int
		days          int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "Validate 18 days",
			args: args{
				initialValues: []int{3, 4, 3, 1, 2},
				days:          18,
			},
			result: 26,
		},
		{
			name: "Validate 80 days",
			args: args{
				initialValues: []int{3, 4, 3, 1, 2},
				days:          80,
			},
			result: 5934,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spendDays(tt.args.initialValues, tt.args.days)
			if len(lanternschool) != tt.result {
				t.Fatalf("Lanternfish count not correct.\nExpected:\n%v\nActual:\n%v\n",
					tt.result,
					len(lanternschool))
			}
			lanternschool = []*model.Lanternfish{}
		})
	}
}
