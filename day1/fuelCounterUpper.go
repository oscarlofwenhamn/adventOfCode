package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Boring stuff
func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}

func splitAndConvertString(str string) []int {
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

// End boring stuff
// Begin funn-er stuff
func recursiveCounter(mass int) int {
	var extraFuel int
	extraFuel = fuelCounter(mass)
	if extraFuel > 0 {
		return extraFuel + recursiveCounter(extraFuel)
	}
	return 0
}

func fuelCounter(mass int) int {
	if mass <= 0 {
		return 0
	}
	fuel := math.Trunc(float64(mass)/3) - 2
	return int(fuel)
}

func main() {
	massString := readFile("./input.txt")
	massList := splitAndConvertString(massString)
	var fuel int
	for _, mass := range massList {
		fuel += recursiveCounter(mass)
	}
	fmt.Println(fuel)
}
