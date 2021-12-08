package day07

import (
	"testing"
)

func TestCost(t *testing.T) {
	positions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	tests := []struct {
		position int
		cost     int
	}{
		{position: 2, cost: 37},
		{position: 1, cost: 41},
		{position: 3, cost: 39},
		{position: 10, cost: 71},
	}
	for i, test := range tests {
		cost := costToMove(positions, test.position)
		if test.cost != cost {
			t.Errorf("for test %d expected %d, got %d", i, test.cost, cost)
		}
	}
}

func TestPart1(t *testing.T) {
	position, cost := part1("example-input.txt")
	if position != 2 || cost != 37 {
		t.Errorf("expected position=2, cost=37, got position=%d, cost=%d", position, cost)
	}
}

func TestPart2(t *testing.T) {
	position, cost := part2("example-input.txt")
	if position != 5 || cost != 168 {
		t.Errorf("expected position=5, cost=168, got position=%d, cost=%d", position, cost)
	}
}
