package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

type vent struct {
	level *int
}

func (v vent) getLevelValue() string {
	switch *v.level {
	case 0:
		return "."
	default:
		return strconv.Itoa(*v.level)
	}
}

func (v vent) increaseLevel() {
	*v.level++
}

type pair struct {
	x int
	y int
}

func convertEntry(entry string) (ret []pair, err error) {

	var indexes = regexp.MustCompile(`\d+`)

	var matches = indexes.FindAllString(entry, -1)

	var xStart, xEnd, yStart, yEnd int

	var results = [4]*int{&xStart, &yStart, &xEnd, &yEnd}
	for i, v := range matches {
		conv, err := strconv.Atoi(v)

		if err != nil {
			return nil, errors.New("entry conversion failed")
		}

		*results[i] = conv
	}

	if xStart > xEnd {
		xStart, xEnd = xEnd, xStart
	}

	if yStart > yEnd {
		yStart, yEnd = yEnd, yStart
	}

	if (yStart != yEnd) && (xStart != xEnd) {
		return
	}

	for i := xStart; i <= xEnd; i++ {
		for j := yStart; j <= yEnd; j++ {
			ret = append(ret, pair{i, j})
		}
	}

	return
}

func getVentLevels(coord []pair) map[int]int {
	coordinateLevels := make(map[pair]int, 0)

	for _, v := range coord {
		coordinateLevels[v]++
	}

	levelMap := make(map[int]int, 0)
	for _, v := range coordinateLevels {
		levelMap[v]++
	}

	return levelMap
}

func main() {
	inputArray := utils.ReadStringArrayFromFile("./data/input")

	entries := []pair{}
	for _, v := range inputArray {
		entry, err := convertEntry(v)
		if err != nil {
			panic("entry conversion failed")
		}
		entries = append(entries, entry...)
	}

	levels := getVentLevels(entries)

	var sum int
	for k, v := range levels {
		if k > 1 {
			sum += v
		}
	}
	fmt.Println(sum)
}
