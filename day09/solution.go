package day09

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
}

func part1(inputpath string) int {
	grid := helpers.FileToStringSlice(inputpath)
	return sumLowPoints(grid)
}

func sumLowPoints(grid []string) int {
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			height := valueAt(grid, x, y)
			leftHeight := valueAt(grid, x-1, y)
			rightHeight := valueAt(grid, x+1, y)
			upHeight := valueAt(grid, x, y-1)
			downHeight := valueAt(grid, x, y+1)
			if height < leftHeight && height < rightHeight && height < upHeight && height < downHeight {
				sum += height + 1
			}
		}
	}
	return sum
}

func valueAt(grid []string, x, y int) int {
	if y < 0 || y >= len(grid) {
		return 10
	}
	if x < 0 || x >= len(grid[y]) {
		return 10
	}
	return int(grid[y][x] - '0')
}
