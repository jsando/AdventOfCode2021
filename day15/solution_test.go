package day15

import (
	"testing"
)

func TestExample(t *testing.T) {
	grid := NewGridFromFile("example-input.txt")
	if grid.Size() != 10 {
		t.Errorf("expected size of 10, got %d", grid.Size())
	}
	if grid.Get(0, 0) != 1 {
		t.Errorf("wrong value")
	}
	if grid.Get(1, 1) != 3 {
		t.Errorf("wrong value")
	}
	if grid.Get(9, 9) != 1 {
		t.Errorf("wrong value")
	}
}

func TestPart1(t *testing.T) {
	cost := part1("example-input.txt")
	if cost != 40 {
		t.Errorf("expected cost of 40, got %d", cost)
	}
}

func TestPart2(t *testing.T) {
	cost := part2("example-input.txt")
	if cost != 315 {
		t.Errorf("expected cost of 315, got %d", cost)
	}
}
