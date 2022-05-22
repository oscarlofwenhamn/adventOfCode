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
	fuelConsumption := alignCrabsMedian(floats)

	fmt.Printf("Total fuel consumption for median: %v\n", fuelConsumption)

	fuelConsumption = alignCrabsMean(floats)
	fmt.Printf("Total fuel consumption for mean: %v\n", fuelConsumption)
}

func alignCrabsMedian(startingPositions []float64) float64 {
	sort.Float64s(startingPositions)

	median := getMedian(startingPositions)
	fmt.Printf("Median: %v\n", median)

	var fuelForMedian float64
	for _, v := range startingPositions {
		fuelForMedian += math.Abs(v - median)
	}

	return fuelForMedian
}

func alignCrabsMean(startingPositions []float64) float64 {
	sort.Float64s(startingPositions)

	mean := getMean(startingPositions)
	fmt.Printf("Mean: %v\n", mean)

	meanFloor := math.Floor(mean)
	var fuelForMeanFloor float64
	for _, v := range startingPositions {
		fuelForMeanFloor += fuel(math.Abs(v - meanFloor))
	}
	meanCeil := math.Ceil(mean)
	var fuelForMeanCeil float64
	for _, v := range startingPositions {
		fuelForMeanCeil += fuel(math.Abs(v - meanCeil))
	}

	if fuelForMeanCeil < fuelForMeanFloor {
		return fuelForMeanCeil
	}
	return fuelForMeanFloor
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

func fuel(n float64) float64 {
	if n == 0 {
		return 0
	}
	return n + fuel(n-1)
}
