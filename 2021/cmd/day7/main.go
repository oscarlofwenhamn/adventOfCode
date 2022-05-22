package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

func main() {
	input := strings.Split(utils.ReadStringArrayFromFile("./input")[0], ",")
	fmt.Println(input)
	floats := []float64{}
	for _, v := range input {
		i, _ := strconv.ParseFloat(v, 64)
		floats = append(floats, i)
	}
	println(len(floats))
	fuelConsumption := alignCrabs(floats)

	fmt.Printf("Total fuel consumption: %v\n", fuelConsumption)
}

func alignCrabs(startingPositions []float64) float64 {
	sort.Float64s(startingPositions)

	median := getMedian(startingPositions)
	fmt.Printf("Median: %v\n", median)

	var fuelForMedian float64
	for _, v := range startingPositions {
		fuelForMedian += math.Abs(v - median)
	}

	return fuelForMedian
}

func getMedian(a []float64) float64 {
	if len(a)%2 == 0 {
		return (a[(len(a)-1)/2] + a[(len(a)+1)/2]) / 2
	}
	return a[len(a)/2]
}

func getMean(a []float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum / float64(len(a))
}
