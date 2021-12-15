package day09

import "testing"

func TestLowPoint(t *testing.T) {
	lines := []string{
		"2199943210",
	}
	sum := sumLowPoints(lines)
	if sum != 3 {
		t.Errorf("expected 3, got %d", sum)
	}
}

func TestPart1(t *testing.T) {
	sum := part1("example-input.txt")
	if sum != 15 {
		t.Errorf("expected 15, got %d", sum)
	}
}
