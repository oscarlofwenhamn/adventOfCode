package main

var scores map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {

}

func getErrorScore(lines []string) int {

	panic("not implemented")
}

func validateChunk() error {
	// if next rune is opening; start new chunk, otherwise validate correct closing symbol or "throw"
	panic("not implemented")
}
