package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/day6/model"
	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

var lanternschool = []*model.Lanternfish{}
var days int = 80

func main() {

	input := utils.ReadStringArrayFromFile("./input")[0]
	integerStrings := strings.Split(input, ",")
	initialValues := []int{}
	for _, v := range integerStrings {
		i, err := strconv.Atoi(v)
		utils.Check(err)
		initialValues = append(initialValues, i)
	}
	spendDays(initialValues, days)
	fmt.Printf("Lanternfish after %v: %v\n", days, len(lanternschool))
}

func spendDays(initialValues []int, days int) {
	for _, value := range initialValues {
		lanternschool = append(lanternschool, &model.Lanternfish{Timer: value})
	}

	for i := 0; i < days; i++ {

		for _, fish := range lanternschool {
			spendDay(fish)
		}

	}
}

func spendDay(fish *model.Lanternfish) {
	if fish.Timer == 0 {
		lanternschool = append(lanternschool, &model.Lanternfish{Timer: 8})
		fish.Timer = 6
	} else {
		fish.Timer--
	}
}
