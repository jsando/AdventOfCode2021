package day07

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	position, cost := part1(inputpath)
	fmt.Printf("Part 1: position=%d, cost=%d\n", position, cost)
	position, cost = part2(inputpath)
	fmt.Printf("Part 2: position=%d, cost=%d\n", position, cost)
}

func part1(inputpath string) (position, cost int) {
	positions := helpers.StringToIntSlice(helpers.FileToStringSlice(inputpath)[0], ",")
	min, max := helpers.MinMaxIntSlice(positions)
	minCost := -1
	minPosition := -1
	for i := min; i <= max; i++ {
		c := costToMove(positions, i)
		if minCost == -1 || c < minCost {
			minCost = c
			minPosition = i
		}
	}
	return minPosition, minCost
}

func costToMove(positions []int, position int) int {
	cost := 0
	for _, val := range positions {
		cost += (helpers.AbsInt(val - position))
	}
	return cost
}

func part2(inputpath string) (position, cost int) {
	positions := helpers.StringToIntSlice(helpers.FileToStringSlice(inputpath)[0], ",")
	min, max := helpers.MinMaxIntSlice(positions)
	minCost := -1
	minPosition := -1
	for i := min; i <= max; i++ {
		c := cost2Move(positions, i)
		if minCost == -1 || c < minCost {
			minCost = c
			minPosition = i
		}
	}
	return minPosition, minCost
}

func cost2Move(positions []int, position int) int {
	cost := 0
	for _, val := range positions {
		diff := (helpers.AbsInt(val - position))
		for diff > 0 {
			cost += diff
			diff--
		}
	}
	return cost
}
