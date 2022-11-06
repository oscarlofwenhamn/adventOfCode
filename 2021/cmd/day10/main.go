package main

import "errors"

var openers []rune = []rune{'(', '[', '{', '<'}
var closers map[rune]int = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

func main() {}

type sequence struct {
	str string
	i   int
}

func (s sequence) peek() (rune, error) {
	if s.i+1 > len(s.str) {
		return 0, errors.New("end of string")
	}
	return rune(s.str[s.i+1]), nil
}

func (s *sequence) pop() (rune, error) {
	if s.i+1 > len(s.str) {
		return 0, errors.New("end of string")
	}
	s.i++
	return rune(s.str[s.i]), nil
}

func newSequence(s string) *sequence {
	return &sequence{s, 0}
}
