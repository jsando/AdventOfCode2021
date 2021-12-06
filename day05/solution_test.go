package day05

import "testing"

func Test_Part1(t *testing.T) {
	count := part1("example-input.txt")
	if count != 5 {
		t.Errorf("expected 5, got %d", count)
	}
}

func Test_Part2(t *testing.T) {
	count := part2("example-input.txt")
	if count != 12 {
		t.Errorf("expected 12, got %d", count)
	}
}
