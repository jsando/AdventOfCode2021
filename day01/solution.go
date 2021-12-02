package day01

import (
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputPath string) {
	depths := helpers.FileToIntSlice(inputPath)
	fmt.Printf("Part 1: %d\n", countIncreases(depths))
	fmt.Printf("Part 2: %d\n", countIncreases(get3WindowSum(depths)))
}

func countIncreases(depths []int) int {
	if len(depths) == 0 {
		return 0
	}
	count := 0
	last := depths[0]
	for i := 1; i < len(depths); i++ {
		if depths[i] > last {
			count++
		}
		last = depths[i]
	}
	return count
}

func get3WindowSum(ints []int) []int {
	sums := make([]int, 0)
	for i := 0; i < len(ints)-2; i++ {
		sums = append(sums, ints[i]+ints[i+1]+ints[i+2])
	}
	return sums
}
