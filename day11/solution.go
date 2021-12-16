package day11

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	flashed := 0
	cavern := NewOctopusCavernFromFile(inputpath)
	for i := 0; i < 100; i++ {
		count, _ := cavern.Step()
		flashed += count
	}
	return flashed
}

func part2(inputpath string) int {
	cavern := NewOctopusCavernFromFile(inputpath)
	step := 1
	for {
		_, allFlashed := cavern.Step()
		if allFlashed {
			return step
		}
		step++
	}
}

type OctopusCavern struct {
	width, height int
	octopuses     [][]int
}

func NewOctopusCavernFromFile(inputpath string) *OctopusCavern {
	lines := helpers.FileToStringSlice(inputpath)
	width := len(lines[0])
	cavern := &OctopusCavern{
		octopuses: make([][]int, 0),
		width:     width,
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if len(line) != width {
			panic(fmt.Sprintf("Expected width %d, got '%s'", width, line))
		}
		row := make([]int, width)
		for x := 0; x < width; x++ {
			row[x] = int(line[x] - '0')
		}
		cavern.octopuses = append(cavern.octopuses, row)
	}
	cavern.height = len(cavern.octopuses)
	return cavern
}

func (c *OctopusCavern) EnergyAt(x, y int) int {
	return c.octopuses[y][x]
}

// Increment the energy level of the given octopus, while
// ignoring if coordinates don't exist or octopuses that
// flashed already.
func (c *OctopusCavern) Increment(x, y int) {
	if x < 0 || x >= c.width {
		return
	}
	if y < 0 || y >= c.height {
		return
	}
	energy := c.EnergyAt(x, y)
	if energy >= 0 {
		c.Set(x, y, energy+1)
	}
}

func (c *OctopusCavern) Set(x, y, level int) {
	c.octopuses[y][x] = level
}

func (c *OctopusCavern) Step() (int, bool) {
	// "First, the energy level of each octopus increases by 1."
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.Increment(x, y)
		}
	}

	// "Then, any octopus with an energy level greater than 9 flashes. This increases
	// the energy level of all adjacent octopuses by 1, including octopuses that are
	// diagonally adjacent. If this causes an octopus to have an energy level
	// greater than 9, it also flashes. This process continues as long as new
	// octopuses keep having their energy level increased beyond 9."
	totalFlashes := 0
	for {
		count := 0
		for y := 0; y < c.height; y++ {
			for x := 0; x < c.width; x++ {
				if c.EnergyAt(x, y) > 9 {
					count++
					c.Set(x, y, -1)     // mark as flashed
					c.Increment(x, y-1) // going clockwise from 12pm
					c.Increment(x+1, y-1)
					c.Increment(x+1, y)
					c.Increment(x+1, y+1)
					c.Increment(x, y+1)
					c.Increment(x-1, y+1)
					c.Increment(x-1, y)
					c.Increment(x-1, y-1)
				}
			}
		}
		if count == 0 {
			break
		}
		totalFlashes += count
	}

	// "Finally, any octopus that flashed during this step has its energy level set to 0,
	// as it used all of its energy to flash."
	flashCount := 0
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if c.EnergyAt(x, y) == -1 {
				c.Set(x, y, 0)
				flashCount++
			}
		}
	}
	return totalFlashes, flashCount == (c.width * c.height)
}
