package day10

import (
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		chunk  string
		expect rune
	}{
		{chunk: "([])", expect: ' '},
		{chunk: "{()()()}", expect: ' '},
		{chunk: "<([{}])>", expect: ' '},
		{chunk: "[<>({}){}[([])<>]]", expect: ' '},
		{chunk: "(((((((((())))))))))", expect: ' '},
		{chunk: "{([(<{}[<>[]}>{[]{[(<()>", expect: '}'},
		{chunk: "[[<[([]))<([[{}[[()]]]", expect: ')'},
	}
	for _, test := range tests {
		got, _ := parseChunk(test.chunk)
		if got != test.expect {
			t.Errorf("for chunk '%s', expected %c got %c", test.chunk, test.expect, got)
		}
	}
}

func TestPart1(t *testing.T) {
	score := part1("example-input.txt")
	if score != 26397 {
		t.Errorf("expected 26397, got %d", score)
	}
}

func TestPart2(t *testing.T) {
	score := part2("example-input.txt")
	if score != 288957 {
		t.Errorf("expected 288957, got %d", score)
	}
}
