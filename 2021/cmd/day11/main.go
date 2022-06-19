package main

import (
	"strconv"
	"strings"

	m "github.com/oscarlofwenhamn/adventOfCode/cmd/day11/models"
)

var directions []m.Coordinates = []m.Coordinates{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

func main() {

}

func fullDumbo(input []string, steps int) (totalFlashes int) {
	var octopi = generateDumbo(input)

	totalFlashes = runDumbo(octopi, steps)

	return
}

func generateDumbo(input []string) (octopi [][]*m.Octopus) {
	for _, row := range input {
		var values = strings.Split(row, "")
		octopi = append(octopi, func() (octopuses []*m.Octopus) {
			for _, valueString := range values {
				var startValue, _ = strconv.Atoi(valueString)
				octopuses = append(octopuses, &m.Octopus{Energy: startValue})
			}
			return
		}())
	}
	return
}

func runDumbo(octopi [][]*m.Octopus, steps int) (totalFlashes int) {
	var flashList []m.Coordinates

	for i := 0; i < steps; i++ {
		for y, row := range octopi {
			for x, octopus := range row {
				octopus.Energy++
				if octopus.Energy > 9 {
					flashList = append(flashList, m.Coordinates{X: x, Y: y})
				}
			}
		}
		for _, tuple := range flashList {
			boostNeighbors(&octopi, tuple, flashList)
			octopi[tuple.Y][tuple.X].Energy = 0
			totalFlashes++
		}
	}
	return
}

func boostNeighbors(octopi *[][]*m.Octopus, coordinates m.Coordinates, flashList []m.Coordinates) {
	for _, dir := range directions {
		var x = coordinates.X + dir.X
		var y = coordinates.Y + dir.Y
		if x < 0 ||
			x > len((*octopi)[coordinates.X])-1 ||
			y < 0 ||
			y > len(*octopi)-1 {
			continue
		}
		if (*octopi)[y][x].Energy == 0 || (*octopi)[y][x].Energy > 9 {
			continue
		}
		(*octopi)[y][x].Energy++
		if (*octopi)[y][x].Energy > 9 {
			flashList = append(flashList, m.Coordinates{X: x, Y: y})
		}
	}
}
