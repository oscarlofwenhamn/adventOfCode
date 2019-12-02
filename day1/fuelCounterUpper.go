package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	// var mass int = 10
	massString := readFile("./input.txt")
	massList := splitString(massString)
	var fuel int
	for _, mass := range massList {
		fuel += fuelCounter(mass)
	}

	// for diff > 0 {
	// 	diff =
	// }
	fmt.Println(fuel)
}

func fuelCounter(mass int) int {
	var fuelTot float64
	if mass == 0 {
		return 0
	}
	fuel := math.Trunc(float64(mass)/3) - 2
	fuelTot += fuel
	return int(fuelTot)
}

func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}

func splitString(str string) []int {
	split := strings.Split(str, "\n")
	array := make([]int, len(split))
	for i, s := range split {
		if len(s) == 0 {
			continue
		}
		j, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		array[i] = j
	}
	return array
}
