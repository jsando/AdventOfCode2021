package day15

import (
	"image"
	"testing"
)

func TestExample(t *testing.T) {
	grid := NewGridFromFile("example-input.txt")
	if grid.Width() != 10 {
		t.Errorf("expected width of 10, got %d", grid.Width())
	}
	if grid.Height() != 10 {
		t.Errorf("expected height of 10, got %d", grid.Height())
	}
	if grid.CostOf(image.Point{0, 0}) != 1 {
		t.Errorf("wrong value")
	}
	if grid.CostOf(image.Point{1, 1}) != 3 {
		t.Errorf("wrong value")
	}
	if grid.CostOf(image.Point{9, 9}) != 1 {
		t.Errorf("wrong value")
	}
}

func TestPart1(t *testing.T) {
	cost := part1("example-input.txt")
	if cost != 40 {
		t.Errorf("expected cost of 40, got %d", cost)
	}
}
