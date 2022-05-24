package main

import (
	"strconv"
	"strings"
)

type location struct {
	x      int
	y      int
	height int
}

type heightMap struct {
	locations [][]location
}

func (loc location) getScore() int {
	return loc.height + 1
}

func createMap(input []string) (hMap heightMap) {
	for y, line := range input {
		rows := strings.Split(line, "")
		for x, hString := range rows {
			height, _ := strconv.Atoi(hString)
			hMap.locations[x][y] = location{x, y, height}
		}
	}
	return
}

func calculateRisk(hMap heightMap) int {
	panic("not implemented!")
}

func main() {

}
