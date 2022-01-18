package day13

import (
	"testing"
)

func TestFold(t *testing.T) {
	tests := []struct {
		x, y         int
		foldX, foldY int
		newX, newY   int
	}{
		{x: 0, y: 14, foldY: 7, newX: 0, newY: 0},
		{x: 2, y: 14, foldY: 7, newX: 2, newY: 0},
		{x: 27, y: 8, foldY: 7, newX: 27, newY: 6},
		{x: 14, y: 0, foldX: 7, newX: 0, newY: 0},
	}
	for i, test := range tests {
		x, y := foldXY(test.x, test.y, test.foldX, test.foldY)
		if x != test.newX || y != test.newY {
			t.Errorf("on test case %d, expected (%d, %d), got (%d, %d)", i, test.newX, test.newY, x, y)
		}
	}
}

func TestLoadPaper(t *testing.T) {
	paper := NewPaperFromFile("example-input.txt")
	if paper.PointsVisible() != 18 {
		t.Errorf("expected 18 points, got %d", paper.PointsVisible())
	}
	if paper.FoldCount() != 2 {
		t.Errorf("expected 2 fold instructions")
	}
}

func TestPart1(t *testing.T) {
	count := part1("example-input.txt")
	if count != 17 {
		t.Errorf("Expected 17, got %d", count)
	}
}

const exampleExpect = `#####
#   #
#   #
#   #
#####
`

func TestPart2(t *testing.T) {
	s := part2("example-input.txt")
	if s != exampleExpect {
		t.Errorf("expected \n%s\n but got \n%s\n", exampleExpect, s)
	}
}
