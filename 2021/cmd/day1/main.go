/*
.:Adent of Code:.
Day 1: Sonar Sweep

Usage:
go run main.go $INPUT_FILE $WINDOW_SIZE
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	windowSize, err := strconv.Atoi(os.Args[2])
	check(err)
	data, err := os.ReadFile(filePath)
	check(err)
	var stringArr = strings.Split(string(data), "\n")
	var intArr []int
	for _, val := range stringArr {
		if val == "" {
			continue
		}
		var intVal, _ = strconv.Atoi(val)
		intArr = append(intArr, intVal)
	}
	result := measureDepthIncreases(intArr, windowSize)
	fmt.Println(result)
}

func biggerThan(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func measureDepthIncreases(data []int, windowSize int) int {
	var count = 0

	previousWindow := 0
	sizeIndex := windowSize - 1
	for i := range data {

		if i < sizeIndex {
			continue
		}

		currentWindow := 0
		// i+1 to ensure the last element of the array is caught, since [x:y] is y-exclusive.
		for _, y := range data[i-sizeIndex : i+1] {
			currentWindow += y
		}

		if i == sizeIndex {
			previousWindow = currentWindow
			continue
		}

		if biggerThan(currentWindow, previousWindow) {
			count++
		}
		previousWindow = currentWindow
	}
	return count
}
