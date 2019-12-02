package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Boring stuff
// Read fileName into byte-array and return string representation,
// keeping an eye out for errors.
func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}

// Convert each element of string array to int and save in new array
func arrayStringToInt(str []string) []int {
	array := make([]int, len(str))
	for i, s := range str {
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

// Recursively figure out total fuel requirement of module and the fuel that follows
func recursiveCounter(mass int) int {
	var extraFuel int
	extraFuel = fuelCounter(mass)
	if extraFuel > 0 {
		return extraFuel + recursiveCounter(extraFuel)
	}
	return 0
}

// Calculate required mass of fuel required based on mass input
func fuelCounter(mass int) int {
	if mass <= 0 {
		return 0
	}
	fuel := math.Trunc(float64(mass)/3) - 2
	return int(fuel)
}

// Read input from file, split string and convert to int array. Calculate and sum up required
// fuel for each module using recursive counter.
func main() {
	massString := readFile("./input.txt")
	massList := arrayStringToInt(strings.Split(massString, "\n"))
	var fuel int
	for _, mass := range massList {
		fuel += recursiveCounter(mass)
	}
	fmt.Println(fuel)
}

// 5099916
