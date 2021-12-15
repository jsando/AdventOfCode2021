package day09

import (
	"reflect"
	"testing"
)

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

func TestPart2(t *testing.T) {
	sum := part2("example-input.txt")
	if sum != 1134 {
		t.Errorf("expected 1134, got %d", sum)
	}
}

func TestSumBasins(t *testing.T) {
	tests := []struct {
		grid   []string
		basins []int
	}{
		{
			grid:   []string{"2199943210"},
			basins: []int{2, 5},
		},
		{
			grid:   []string{"2199943210", "3987894921"},
			basins: []int{3, 8, 3},
		},
	}
	for _, test := range tests {
		got := findBasins(test.grid)
		if !reflect.DeepEqual(got, test.basins) {
			t.Errorf("expected %v, got %v", test.basins, got)
		}
	}
}
