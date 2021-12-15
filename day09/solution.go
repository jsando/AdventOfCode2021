package day09

import (
	"fmt"
	"sort"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	grid := helpers.FileToStringSlice(inputpath)
	return sumLowPoints(grid)
}

func part2(inputpath string) int {
	input := helpers.FileToStringSlice(inputpath)
	basins := findBasins(input)
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	sum := 1
	for i := 0; i < 3; i++ {
		if i >= len(basins) {
			break
		}
		sum *= basins[i]
	}
	return sum
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

func findBasins(stringGrid []string) []int {
	basins := []int{}
	width := len(stringGrid[0])
	var grid [][]int
	for y := 0; y < len(stringGrid); y++ {
		row := make([]int, width)
		for x := 0; x < width; x++ {
			row[x] = int(stringGrid[y][x] - '0')
		}
		grid = append(grid, row)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < width; x++ {
			count := markRun(grid, x, y)
			if count > 0 {
				basins = append(basins, count)
			}
		}
	}

	return basins
}

func markRun(grid [][]int, x, y int) int {
	count := 0
	if valueAt2(grid, x, y) < 9 {
		grid[y][x] = 9
		count = 1
		count += markRun(grid, x-1, y)
		count += markRun(grid, x+1, y)
		count += markRun(grid, x, y+1)
		count += markRun(grid, x, y-1)
	}
	return count
}

func valueAt2(grid [][]int, x, y int) int {
	if y < 0 || y >= len(grid) {
		return 10
	}
	if x < 0 || x >= len(grid[y]) {
		return 10
	}
	return grid[y][x]
}
