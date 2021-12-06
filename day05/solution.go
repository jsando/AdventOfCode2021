package day05

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	grid := NewGridFromFile(inputpath, false)
	return grid.CountOverlap(2)
}

func part2(inputpath string) int {
	grid := NewGridFromFile(inputpath, true)
	return grid.CountOverlap(2)
}

type Grid struct {
	cells [][]int // [y][x]
}

const gridSize = 1000

func NewGridFromFile(inputpath string, diagonal bool) *Grid {
	grid := &Grid{
		cells: make([][]int, gridSize),
	}
	for y := 0; y < gridSize; y++ {
		grid.cells[y] = make([]int, 1000)
	}

	lines := helpers.FileToStringSlice(inputpath)
	for _, line := range lines {
		var x1, y1, x2, y2 int
		n, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if n != 4 {
			panic(fmt.Sprintf("read %d instead of 4", n))
		}
		if err != nil {
			panic(err)
		}

		if x1 == x2 {
			// vertical
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				grid.cells[y][x1]++
			}
		} else if y1 == y2 {
			// horizontal
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				grid.cells[y1][x]++
			}
		} else if diagonal {
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}
			y := y1
			dy := 1
			if y2 < y1 {
				dy = -1
			}
			for x := x1; x <= x2; x++ {
				grid.cells[y][x]++
				y += dy
			}
		}
	}
	return grid
}

func (g *Grid) CountOverlap(minOverlap int) int {
	count := 0
	for y := 0; y < len(g.cells); y++ {
		for x := 0; x < len(g.cells[y]); x++ {
			if g.cells[y][x] >= minOverlap {
				count++
			}
		}
	}
	return count
}
