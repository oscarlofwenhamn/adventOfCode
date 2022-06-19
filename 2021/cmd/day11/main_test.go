package main

import (
	"testing"

	m "github.com/oscarlofwenhamn/adventOfCode/cmd/day11/models"
)

func Test_Generator(t *testing.T) {

	type args struct {
		input []string
		steps int
	}
	tests := []struct {
		name   string
		args   args
		result [][]m.Octopus
	}{
		{
			name: "Validate test input",
			args: args{
				input: []string{"12345",
					"67890"},
				steps: 10,
			},
			result: [][]m.Octopus{
				{
					{Energy: 1},
					{Energy: 2},
					{Energy: 3},
					{Energy: 4},
					{Energy: 5},
				},
				{
					{Energy: 6},
					{Energy: 7},
					{Energy: 8},
					{Energy: 9},
					{Energy: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var matrix = generateDumbo(tt.args.input)
			for y, row := range matrix {
				for x, value := range row {
					if *value != tt.result[y][x] {
						t.Fatalf("Octopus incorrect!\nExpected: %d\nActual: %d\n", tt.result[y][x], *value)
					}
				}
			}
		})
	}
}

func Test_Booster(t *testing.T) {

	type args struct {
		octopi      [][]*m.Octopus
		coordinates m.Coordinates
		flashList   []m.Coordinates
	}
	type results struct {
		octopi    [][]*m.Octopus
		flashList []m.Coordinates
	}
	tests := []struct {
		name    string
		args    args
		results results
	}{
		{
			name: "Validate base case",
			args: args{
				octopi: [][]*m.Octopus{
					{
						{Energy: 1},
						{Energy: 1},
						{Energy: 1},
					},
					{
						{Energy: 1},
						{Energy: 1},
						{Energy: 1},
					},
					{
						{Energy: 1},
						{Energy: 1},
						{Energy: 1},
					},
				},
				coordinates: m.Coordinates{X: 1, Y: 1},
				flashList:   []m.Coordinates{},
			},
			results: results{
				octopi:    [][]*m.Octopus{{{Energy: 2}, {Energy: 2}, {Energy: 2}}, {{Energy: 2}, {Energy: 1}, {Energy: 2}}, {{Energy: 2}, {Energy: 2}, {Energy: 2}}},
				flashList: []m.Coordinates{},
			},
		},
		{
			name: "Validate flashes added case",
			args: args{
				octopi: [][]*m.Octopus{
					{{Energy: 9}, {Energy: 1}, {Energy: 9}},
					{{Energy: 1}, {Energy: 1}, {Energy: 1}},
					{{Energy: 9}, {Energy: 1}, {Energy: 9}},
				},
				coordinates: m.Coordinates{X: 1, Y: 1},
				flashList:   []m.Coordinates{},
			},
			results: results{
				octopi: [][]*m.Octopus{
					{{Energy: 10}, {Energy: 2}, {Energy: 10}},
					{{Energy: 2}, {Energy: 1}, {Energy: 2}},
					{{Energy: 10}, {Energy: 2}, {Energy: 10}}},
				flashList: []m.Coordinates{{X: 0, Y: 0}, {X: 0, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var foobar = tt.args.octopi

			boostNeighbors(&foobar, tt.args.coordinates, tt.args.flashList)

			for y, row := range foobar {
				for x, value := range row {
					if *value != *tt.results.octopi[y][x] {
						t.Fatalf("Octopus incorrect at (%d,%d)!\nExpected: %d\nActual: %d\n", x, y, *tt.results.octopi[y][x], *value)
					}
					if value.Energy > 9 {
						// verify is in flashList
					}
				}
			}
		})
	}
}
func Test_Full(t *testing.T) {

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
