package day10

import (
	"fmt"
	"sort"
	"strings"
	"text/scanner"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	chunks := helpers.FileToStringSlice(inputpath)
	score := 0
	for _, chunk := range chunks {
		if len(chunk) == 0 {
			continue
		}
		// part 1 only considers errors, ignoring if found == ' '
		found, _ := parseChunk(chunk)
		switch found {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}
	return score
}

func part2(inputpath string) int {
	chunks := helpers.FileToStringSlice(inputpath)
	scores := []int{}
	for _, chunk := range chunks {
		if len(chunk) == 0 {
			continue
		}
		badTok, stack := parseChunk(chunk)
		if badTok == ' ' && !stack.Empty() {
			// Calculate score to complete the expression
			score := 0
			for !stack.Empty() {
				score *= 5
				switch stack.Pop() {
				case ')':
					score += 1
				case ']':
					score += 2
				case '}':
					score += 3
				case '>':
					score += 4
				}
			}
			scores = append(scores, score)
			fmt.Printf("chunk '%s', completion score %d\n", chunk, score)
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

func parseChunk(chunk string) (rune, *RuneStack) {
	var s scanner.Scanner
	s.Init(strings.NewReader(chunk))
	stack := NewRuneStack()
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '(':
			stack.Push(')')
		case '<':
			stack.Push('>')
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		case ')', '>', '}', ']':
			expect := stack.Pop()
			if tok != expect {
				return tok, stack
			}
		default:
			panic(fmt.Sprintf("bad token: %c", tok))
		}
	}
	return ' ', stack
}

type RuneStack struct {
	stack []rune
}

func NewRuneStack() *RuneStack {
	return &RuneStack{stack: make([]rune, 0)}
}

func (s *RuneStack) Empty() bool {
	return len(s.stack) == 0
}

func (s *RuneStack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *RuneStack) Pop() rune {
	if len(s.stack) == 0 {
		panic("empty stack")
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}
