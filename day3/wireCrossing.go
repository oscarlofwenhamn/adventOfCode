package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}

func draw(arr []string) map[[2]int]int {
	x := 0
	y := 0
	wireMap := make(map[[2]int]int)
	fmt.Printf("x: %d, y: %d\n", x, y)

	for _, full := range arr {
		fmt.Println(full)
		steps, _ := strconv.Atoi(strings.TrimRight(full[1:], "\r\n"))
		switch direction := full[0]; direction {
		case 'U':
			fmt.Println(steps)
			for i := 1; i <= steps; i++ {
				coord := [2]int{x, y + i}
				fmt.Println(coord)
				wireMap[coord] = 1
			}
			y += steps
		case 'D':
			fmt.Println(steps)
			for i := 1; i <= steps; i++ {
				coord := [2]int{x, y - i}
				fmt.Println(coord)
				wireMap[coord] = 1
			}
			y -= steps
		case 'R':
			fmt.Println(steps)
			for i := 1; i <= steps; i++ {
				coord := [2]int{x + i, y}
				fmt.Println(coord)
				wireMap[coord] = 1
			}
			x += steps
		case 'L':
			fmt.Println(steps)
			for i := 1; i <= steps; i++ {
				coord := [2]int{x - i, y}
				fmt.Println(coord)
				wireMap[coord] = 1
			}
			x -= steps
		}
		fmt.Printf("x: %d, y: %d\n", x, y)

	}

	return wireMap
}

func getDist(coord [2]int) int {
	return int(math.Abs(float64(coord[0] + coord[1])))
}

func getMin(arr []int) int {
	dist := 0
	for i, num := range arr {
		if i == 0 || num < dist {
			dist = num
		}
	}
	return dist
}

func main() {
	input := readFile("input1.txt")
	wires := strings.Split(input, "\n")
	fmt.Println(wires[0] + "\n" + wires[1])

	wireMap0 := draw(strings.Split(wires[0], ","))
	wireMap1 := draw(strings.Split(wires[1], ","))
	var crosses [][2]int
	var dists []int

	for key := range wireMap0 {
		if wireMap1[key] == 1 && key != [2]int{0, 0} {
			crosses = append(crosses, key)
			dists = append(dists, getDist(key))
		}
	}
	minDist := getMin(dists)
	fmt.Printf("Min dist: %d", minDist)
}
