/*
.:Adent of Code:.
Day 2: Dive!

Usage:
go run main.go $INPUT_FILE
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
	data, err := os.ReadFile(filePath)
	check(err)
	var commandArray = strings.Split(string(data), "\n")

	horizontal, depth, aim := calculatePosition(commandArray, 0, 0, 0)

	fmt.Printf("Horizontal: %d, Depth: %d, Aim: %d\n", horizontal, depth, aim)
	fmt.Printf("Multiple: %d\n", horizontal*depth)

}

func calculatePosition(commands []string, horizontal, depth, aim int) (int, int, int) {
	for _, val := range commands {
		if val == "" {
			continue
		}
		commandRaw := strings.Split(val, " ")
		command := commandRaw[0]
		magnitude, err := strconv.Atoi(commandRaw[1])
		if err != nil {
			continue
		}
		switch command {
		case "forward":
			moveForward(&horizontal, &depth, &aim, magnitude)
		case "down":
			moveDown(&aim, magnitude)
		case "up":
			moveUp(&aim, magnitude)
		}
	}
	return horizontal, depth, aim
}

func moveForward(horizontal, depth, aim *int, val int) {
	*horizontal += val
	*depth += *aim * val
}
func moveDown(aim *int, val int) {
	*aim += val
}
func moveUp(aim *int, val int) {
	*aim -= val
}
