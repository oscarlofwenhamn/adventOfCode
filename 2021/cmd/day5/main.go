package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

	var indexes = regexp.MustCompile(`\d`)

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

	for i := xStart; i <= xEnd; i++ {
		for j := yStart; j <= yEnd; j++ {
			ret = append(ret, pair{i, j})
		}
	}

	return
}

func drawDiagram() [][]vent {
	return [][]vent{}
}

func main() {
	fmt.Println("Hello, World!")
}
