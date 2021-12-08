package day06

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	school := helpers.StringToIntSlice(helpers.FileToStringSlice(inputpath)[0], ",")
	for i := 0; i < 80; i++ {
		school = schoolTick(school)
	}
	return len(school)
}

func part2(inputpath string) int {
	// Convert from array of individual fish ages to array of count by age
	fishAges := helpers.StringToIntSlice(helpers.FileToStringSlice(inputpath)[0], ",")
	school := make([]int, 9)
	for _, fish := range fishAges {
		school[fish]++
	}

	// Simulate 256 days of life
	for i := 0; i < 256; i++ {
		school2Tick(school)
	}

	// Return count of how many fish after 256 days
	sum := 0
	for _, fishCount := range school {
		sum += fishCount
	}
	return sum
}

func fishTick(timerIn int) (int, []int) {
	if timerIn == 0 {
		return 6, []int{8}
	}
	return timerIn - 1, []int{}
}

func schoolTick(school []int) []int {
	schoolSpawn := []int{}
	for i, timer := range school {
		newTimer, spawn := fishTick(timer)
		school[i] = newTimer
		schoolSpawn = append(schoolSpawn, spawn...)
	}
	return append(school, schoolSpawn...)
}

// Instead of a list of fish, track how many fish are in each state so it needs
// at most 9 ints instead of billions. each tick shifts the array left, adding
// the count from the leftmost element to both element 6 (to reset their spwawn
// count) and element 8 (spawning new fish).
func school2Tick(school []int) {
	spawnCount := school[0]
	for i := 1; i < 9; i++ {
		school[i-1] = school[i]
	}
	school[8] = spawnCount
	school[6] += spawnCount
}
