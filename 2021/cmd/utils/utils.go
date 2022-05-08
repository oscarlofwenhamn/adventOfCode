package utils

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadStringArrayFromFile(path string) []string {
	input, err := os.ReadFile(path)
	check(err)
	return strings.Split(string(input), "\n")
}
