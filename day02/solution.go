package day02

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

// Run run reindeer.
func Run(inputpath string) {
	moves := helpers.FileToStringSlice(inputpath)
	fmt.Printf("Part 1: %d\n", part1(moves))
	fmt.Printf("Part 2: %d\n", part2(moves))
}

func part1(moves []string) int {
	var sub Submarine
	sub.MoveSeries(moves)
	return sub.Horizontal * sub.Depth
}

func part2(moves []string) int {
	var sub Submarine2
	for _, move := range moves {
		sub.Move(move)
	}
	return sub.Horizontal * sub.Depth
}

type Submarine struct {
	Horizontal int
	Depth      int
}

func (s *Submarine) Move(spec string) {
	var direction string
	var distance int
	n, err := fmt.Sscanf(spec, "%s %d", &direction, &distance)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("exptected 2 args")
	}
	switch direction {
	case "forward":
		s.Horizontal += distance
	case "up":
		s.Depth -= distance
	case "down":
		s.Depth += distance
	default:
		panic(direction)
	}
}

func (s *Submarine) MoveSeries(moves []string) {
	for _, move := range moves {
		s.Move(move)
	}
}

type Submarine2 struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (s *Submarine2) Move(spec string) {
	var direction string
	var distance int
	n, err := fmt.Sscanf(spec, "%s %d", &direction, &distance)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("exptected 2 args")
	}
	switch direction {
	case "forward":
		s.Horizontal += distance
		s.Depth += s.Aim * distance
	case "up":
		s.Aim -= distance
	case "down":
		s.Aim += distance
	default:
		panic(direction)
	}
}
