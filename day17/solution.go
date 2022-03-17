package day17

import (
	"fmt"
)

func Run(inputPath string) {
	//fmt.Printf("Part 1: %d\n", part1(inputPath))
}

func part1(inputPath string) int {
	//lines := helpers.FileToStringSlice(inputPath)
	//x1, y1, x2, y2 := parseTarget(lines[0])
	//if y2 > y1 {
	//	y1, y2 = y2, y1
	//}
	//maxY := 0
	//minX := findMinSum(x1) // minimum horizontal velocity that can possibly intersect target
	//maxX := x2             // maximum horizontal velocity that can still intersect target
	//maxSteps := minX
	//for i := maxX; i <= maxX; i++ {
	//
	//}
	//for dx := minX; dx <= maxX; dx++ {
	//	// haha ok can't just test "in range", need to know maxY achieved
	//	// would also help if we knew that it failed because x never made it to target,
	//	// because then any higher values of Y can be discarded
	//	for dy := y2; dy <=
	//	if inRange(x1, y1, x2, y2, dx, dy) {
	//
	//	}
	//
	//}
	//return maxY
	return 0
}

func findMinSum(target int) int {
	i := 0
	sum := 0
	for sum < target {
		i++
		sum += i
	}
	return i
}

func parseTarget(target string) (x1, y1, x2, y2 int) {
	n, err := fmt.Sscanf(target, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		panic("expected to parse 4 params")
	}
	return
}

type Probe struct {
	x, y   int
	dx, dy int
}

func computeStep(probe Probe) Probe {
	dx := probe.dx
	if dx < 0 {
		dx++
	}
	if dx > 0 {
		dx--
	}
	dy := probe.dy - 1
	return Probe{probe.x + probe.dx, probe.y + probe.dy, dx, dy}
}

func inRange(x1, y1, x2, y2, dx, dy int) bool {
	// I prefer to think of y1 as the top and y2 as the bottom
	if y2 > y1 {
		y1, y2 = y2, y1
	}
	probe := Probe{x: 0, y: 0, dx: dx, dy: dy}
	for {
		probe = computeStep(probe)
		// Inside target range?
		if probe.x >= x1 && probe.x <= x2 && probe.y <= y1 && probe.y >= y2 {
			return true
		}
		// If left of range, and dx <= 0 then we are no longer closing and can abort
		if probe.x < x1 && probe.dx <= 0 {
			return false
		}
		// If right of range and dx >= 0 we can abort
		if probe.x > x2 && probe.dx >= 0 {
			return false
		}
		// If below range and dy <= 0 we can abort
		if probe.y < y2 && probe.dy <= 0 {
			return false
		}
	}
}
