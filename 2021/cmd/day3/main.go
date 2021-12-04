/*
.:Adent of Code:.
Day 3: Binary Diagnostic

Usage:
go run main.go $INPUT_FILE
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const numLen = 12

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	check(err)
	stringArray := strings.Split(string(data), "\n")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		bitCount := getBitCount(stringArray)
		calculatePowerConsumption(bitCount)
	}()

	var oxygenValue int64
	wg.Add(1)
	go func(oxygenValue *int64) {
		defer wg.Done()
		*oxygenValue = getOxygenValue(stringArray)
	}(&oxygenValue)
	var CO2Value int64

	wg.Add(1)
	go func(CO2Value *int64) {
		defer wg.Done()
		*CO2Value = getCO2Value(stringArray)
	}(&CO2Value)

	wg.Wait()

	fmt.Printf("Life support rating: %d\n", oxygenValue*CO2Value)
}

func getOxygenValue(stringArray []string) int64 {
	getTarget := func(i int) int {
		if i >= 0 {
			return 1
		} else {
			return 0
		}
	}
	oxygenValue := getValue(stringArray, getTarget, 0)
	oxygenValueInt, _ := strconv.ParseInt(oxygenValue, 2, 64)
	fmt.Printf("Oxygen Generator rating: %d\n", oxygenValueInt)
	return oxygenValueInt
}
func getCO2Value(stringArray []string) int64 {
	getTarget := func(i int) int {
		if i < 0 {
			return 1
		} else {
			return 0
		}
	}
	CO2Value := getValue(stringArray, getTarget, 0)
	CO2ValueInt, _ := strconv.ParseInt(CO2Value, 2, 64)
	fmt.Printf("CO2 Scrubber rating: %d\n", CO2ValueInt)
	return CO2ValueInt
}

func getValue(values []string, getTarget func(int) int, index int) string {
	var remainingValues []string
	bitCountValue := getBitCount(values)[index]
	target := getTarget(bitCountValue)

	for _, val := range values {

		if len(val) == 0 {
			continue
		}
		bit, _ := strconv.Atoi(string(val[index]))
		if bit == target {
			remainingValues = append(remainingValues, val)
		}
	}
	if len(remainingValues) == 1 {
		return remainingValues[0]
	} else {
		return getValue(remainingValues, getTarget, index+1)
	}
}

func getBitCount(stringArray []string) []int {
	bitCount := make([]int, numLen)

	for _, number := range stringArray {
		for i, bit := range number {
			switch bit {
			case '1':
				bitCount[i]++
			case '0':
				bitCount[i]--
			}
		}
	}
	return bitCount
}

func calculatePowerConsumption(bitCount []int) {
	var gamma string
	var epsilon string
	for _, count := range bitCount {
		if count > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	fmt.Printf("Gamma rate: %s\nEpsilon rate: %s\n", gamma, epsilon)

	gammaVal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 64)
	powerConsumption := gammaVal * epsilonVal
	fmt.Printf("Power consumption: %d\n", powerConsumption)
}
