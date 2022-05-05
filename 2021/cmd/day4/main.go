/*
.:Adent of Code:.
Day 4: Giant Squid

Usage:
go run main.go $INPUT_FILE

File reader functionality heavily inspired by SO user
Malcom: https://stackoverflow.com/a/12206584/8149467
*/
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BingoCard struct {
	columnMarks [width]int
	rowMarks    [height]int
	numbers     map[int][4]int
	card        [height][width]int
}

const (
	width  = 5
	height = 5
)

func newBingoCard(rows [height][width]int) *BingoCard {
	card := new(BingoCard)
	card.numbers = make(map[int][4]int)

	for i, row := range rows {
		for y, val := range row {
			card.numbers[val] = [4]int{i, y, 1}
			card.card[i][y] = val
		}
	}
	return card
}

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

	var matrix [5][5]int
	var bingoCards []*BingoCard
	for i, line := range stringArray[2:] {
		if line == "" {
			bingoCard := newBingoCard(matrix)

			bingoCards = append(bingoCards, bingoCard)
			continue
		}

		re := regexp.MustCompile("[0-9]+")
		stringArr := re.FindAllString(line, -1)
		numArr := func(row []string) [5]int {
			var arr [5]int

			for i, e := range row {
				intE, _ := strconv.Atoi(e)
				arr[i] = intE
			}
			return arr
		}(stringArr)
		matrix[i%6] = numArr
	}

	bingo := false
	for _, numberString := range strings.Split(stringArray[0], ",") {
		number, _ := strconv.Atoi(numberString)
		for i, card := range bingoCards {
			if card.numbers[number][2] == 1 {
				x := card.numbers[number][0]
				y := card.numbers[number][1]
				card.columnMarks[y]++
				card.rowMarks[x]++
				arr := card.numbers[number]
				arr[3] = 1
				card.numbers[number] = arr

				if card.rowMarks[x] == 5 {
					fmt.Printf("Bingo on card %d, row %d!\n", i+1, y+1)
					calculateWinningScore(card, card.card[x], number)
					bingo = true
					break
				}
				if card.columnMarks[y] == 5 {
					fmt.Printf("Bingo on card %d, column %d!\n", i+1, x+1)
					var winner [5]int
					for i := 0; i < 5; i++ {
						winner[i] = card.card[i][y]
					}
					calculateWinningScore(card, winner, number)
					bingo = true
					break
				}
			}
		}
		if bingo {
			break
		}
	}
}

func calculateWinningScore(card *BingoCard, winningNumbers [5]int, lastNumberCalled int) {
	var sum int
	fmt.Println("Winning numbers:")
	fmt.Println(winningNumbers)
	fmt.Printf("Final number: %d\n", lastNumberCalled)
	for val, arr := range card.numbers {
		if arr[3] != 1 {
			sum += val
		}
	}
	product := sum * lastNumberCalled
	fmt.Printf("Final score: %d * %d = %d\n", sum, lastNumberCalled, product)
}
