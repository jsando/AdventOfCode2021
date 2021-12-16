package main

import (
	"flag"

	"github.com/jsando/AdventOfCode2021/day01"
	"github.com/jsando/AdventOfCode2021/day02"
	"github.com/jsando/AdventOfCode2021/day03"
	"github.com/jsando/AdventOfCode2021/day04"
	"github.com/jsando/AdventOfCode2021/day05"
	"github.com/jsando/AdventOfCode2021/day06"
	"github.com/jsando/AdventOfCode2021/day07"
	"github.com/jsando/AdventOfCode2021/day08"
	"github.com/jsando/AdventOfCode2021/day09"
	"github.com/jsando/AdventOfCode2021/day10"
)

var day = flag.Int("d", 0, "day number (1...25)")
var inputPath = flag.String("i", "", "optional input filename")

type runner func(inputPath string)

var runners []runner = []runner{
	day01.Run, day02.Run, day03.Run, day04.Run, day05.Run,
	day06.Run, day07.Run, day08.Run, day09.Run, day10.Run,
	//day11.Run, day12.Run, day13.Run, day14.Run, day15.Run,
	//day16.Run, day17.Run, day18.Run, day19.Run, day20.Run,
	//day21.Run, day22.Run, day23.Run, day24.Run, day25.Run,
}

func main() {
	flag.Parse()
	runners[*day-1](*inputPath)
}
