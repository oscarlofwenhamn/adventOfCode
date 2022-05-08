package main

import (
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

func main() {

	inputArray := utils.ReadStringArrayFromFile("data/input")

	createDiagram(inputArray)
}

func createDiagram(input []string) []string {

	matrix := [10][10]int{}

	for _, line := range input {
		//0,9 -> 5,9
		matrix = convertSegment(line, matrix)
	}

	return input
}

func convertSegment(segment string, matrix [10][10]int) [10][10]int {
	coordinates := strings.Split(segment, " -> ")
	startX, startY := func(arr []string) (int, int) {
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		return x, y
	}(strings.Split(coordinates[0], ","))
	endX, endY := func(arr []string) (int, int) {
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		return x, y
	}(strings.Split(coordinates[1], ","))

	switch {
	case startY == endY:
		println("Horizontal line")
		dist := endX - startX
		for i := 0; i < dist; i++ {
			matrix[startX+i][startY]++
		}
	case startX == endX:
		println("Vertical line")
		dist := endY - startY
		for i := 0; i < dist; i++ {
			matrix[startX][startY+i]++
		}
	}
	return matrix
}
