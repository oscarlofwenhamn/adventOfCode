package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

const DAYS, TIMER_SIZE, ADULT_TIMER_RESET int = 256, 9, 6

func main() {

	days := flag.Int("days", DAYS, "Number of days to simulate")
	flag.Parse()

	input := utils.ReadStringArrayFromFile("./input")[0]
	integerStrings := strings.Split(input, ",")
	initialValues := []int{}
	for _, v := range integerStrings {
		i, err := strconv.Atoi(v)
		utils.Check(err)
		initialValues = append(initialValues, i)
	}

	tot := simulateLanternFish(initialValues, *days)

	fmt.Printf("Total number of fish: %v\n", tot)
}

func simulateLanternFish(initialValues []int, days int) int {
	fishArray := [TIMER_SIZE]int{}

	for _, v := range initialValues {
		fishArray[v]++
	}

	for i := 0; i < days; i++ {
		spendDay(&fishArray)
	}

	var tot int
	for _, v := range fishArray {
		tot += v
	}
	return tot
}

func spendDay(arr *[TIMER_SIZE]int) {
	curr, slice := arr[0], arr[1:]
	slice[6] += curr
	slice = append(slice, curr)

	copy(arr[:], slice[:TIMER_SIZE])
}
