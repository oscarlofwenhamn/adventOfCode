package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

type location struct {
	height int
}

type heightMap struct {
	locations [][]location
}

func (loc location) getScore() int {
	return loc.height + 1
}

func (loc location) isHigherThanOrEqualTo(i int) bool {
	return loc.height >= i
}

func createMap(input []string) (hMap heightMap) {

	hMap.locations = make([][]location, len(input))
	for y, line := range input {
		rows := strings.Split(line, "")
		for _, hString := range rows {
			height, _ := strconv.Atoi(hString)
			hMap.locations[y] = append(hMap.locations[y], location{height})
		}
	}
	return
}

func calculateRisk(hMap heightMap) int {
	var risk int

	for y, line := range hMap.locations {
		for x, location := range line {
			if y != 0 && location.isHigherThanOrEqualTo(hMap.locations[y-1][x].height) {
				continue
			}
			if y != len(hMap.locations)-1 && location.isHigherThanOrEqualTo(hMap.locations[y+1][x].height) {
				continue
			}
			if x != 0 && location.isHigherThanOrEqualTo(hMap.locations[y][x-1].height) {
				continue
			}
			if x != len(hMap.locations[0])-1 && location.isHigherThanOrEqualTo(hMap.locations[y][x+1].height) {
				continue
			}
			risk += location.getScore()
		}
	}
	return risk
}

func main() {
	input := utils.ReadStringArrayFromFile("./input")
	hMap := createMap(input)
	risk := calculateRisk(hMap)
	fmt.Println(risk)
}
