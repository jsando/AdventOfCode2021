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
	path := grid.Path(image.Point{X: 0, Y: 0}, image.Point{X: 9, Y: 9})
	grid.PrintPath(path)
}
