package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/day9/models"
	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

func createMap(input []string) (hMap models.HeightMap) {

	hMap.Locations = make([][]models.Location, len(input))
	for y, line := range input {
		rows := strings.Split(line, "")
		for x, hString := range rows {
			height, _ := strconv.Atoi(hString)
			hMap.Locations[y] = append(hMap.Locations[y], models.MakeLocation(height, x, y))
		}
	}

	generateLowPoints(&hMap)
	return
}

func calculateRisk(hMap models.HeightMap) int {
	var risk int

	for _, point := range hMap.LowPoints {
		risk += point.GetScore()
	}
	return risk
}

func generateLowPoints(hMap *models.HeightMap) {
	for y, line := range hMap.Locations {
	Line:
		for x, location := range line {
			for _, neighbour := range location.GetNeighbours(*hMap) {

				if location.IsHigherThanOrEqualTo(neighbour.GetHeight()) {
					continue Line
				}
			}
			hMap.LowPoints = append(hMap.LowPoints, &hMap.Locations[y][x])
		}
	}
}

func calculateBasinValue(hMap models.HeightMap) int {
	var basinSizes []int
	for _, point := range hMap.LowPoints {
		basinSizes = append(basinSizes, getBasinSize(point, &hMap))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	if len(basinSizes) < 3 {
		panic("Too few basins, initial premise broken!")
	}
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func getBasinSize(point *models.Location, hMap *models.HeightMap) int {
	var key [][]bool = make([][]bool, len(hMap.Locations))
	for i := range key {
		key[i] = make([]bool, len(hMap.Locations[i]))
	}
	return countNeighbours(point, hMap, key)
}

func countNeighbours(point *models.Location, hMap *models.HeightMap, key [][]bool) (neighbours int) {
	key[point.GetY()][point.GetX()] = true
	for _, n := range point.GetNeighbours(*hMap) {
		if !key[n.GetY()][n.GetX()] && n.GetHeight() != 9 {
			neighbours += countNeighbours(n, hMap, key)
		}
	}
	return 1 + neighbours
}

func main() {
	input := utils.ReadStringArrayFromFile("./input")
	hMap := createMap(input)
	risk := calculateRisk(hMap)
	basinValue := calculateBasinValue(hMap)
	fmt.Println(risk)
	fmt.Println(basinValue)
}
