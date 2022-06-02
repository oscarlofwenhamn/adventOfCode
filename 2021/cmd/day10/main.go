package main

var scores map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var openers []rune = []rune{'(', '[', '{', '<'}

func main() {

}

func getErrorScore(lines []string) (score int) {

	for _, line := range lines {
		err := validateChunk([]rune(line))
		if err != -1 {
			score += scores[err]
		}
	}
	return
}

func validateChunk(chunk []rune) rune {

	var err rune
	opener := chunk[0]
	chunk = chunk[1:]
	if isOpener(chunk[0]) {
		err = validateChunk(chunk)
	}
	if err != -1 {
		return err
	}
	if isInvalidCloser(opener, chunk[0]) {
		return chunk[0]
	}
	// if next rune is opening; start new chunk, otherwise validate correct closing symbol or "throw"
	return -1
}

func isOpener(str rune) bool {

	for _, o := range openers {
		if str == o {
			return true
		}
	}
	return false
}

func isInvalidCloser(o rune, c rune) bool {
	switch o {
	case '(':
		return c != ')'
	case '[':
		return c != ']'
	case '{':
		return c != '}'
	case '<':
		return c != '>'
	}
	panic("no valid opener supplied")
}
