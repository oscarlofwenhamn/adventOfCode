package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadStringArrayFromFile(path string) []string {
	input, err := os.ReadFile(path)
	Check(err)
	return strings.Split(string(input), "\n")
}
