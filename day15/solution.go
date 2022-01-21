package day15

import (
	"container/heap"
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
	"image"
	"strings"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	grid := NewGridFromFile(inputpath)
	path := grid.Path(image.Point{X: 0, Y: 0}, image.Point{X: grid.Size() - 1, Y: grid.Size() - 1})
	return grid.PathCost(path) - 1 // minus 1 for start
}

func part2(inputpath string) int {
	grid := NewGridFromFile(inputpath)
	grid = grid.SuperSize()
	path := grid.Path(image.Point{X: 0, Y: 0}, image.Point{X: grid.Size() - 1, Y: grid.Size() - 1})
	return grid.PathCost(path) - 1 // minus 1 for start
}

type Grid struct {
	cells []uint8 // cells of grid stored in row-major order
	size  int     // width and height
}

func (r *Grid) Size() int {
	return r.size
}

func (r *Grid) Set(x int, y int, value uint8) {
	r.cells[y*r.size+x] = value
}

func (r *Grid) Get(x int, y int) uint8 {
	return r.cells[y*r.size+x]
}

func NewGrid(size int) *Grid {
	grid := &Grid{
		cells: make([]uint8, size*size),
		size:  size,
	}
	return grid
}

func NewGridFromFile(inputpath string) *Grid {
	// Use length of first line to set grid size, assumes width == height
	lines := helpers.FileToStringSlice(inputpath)
	size := len(strings.TrimSpace(lines[0]))
	grid := NewGrid(size)
	y := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		for x, ch := range line {
			grid.Set(x, y, uint8(ch-'0'))
		}
		y++
	}
	return grid
}

// Path uses A* to find weighted path from start to goal and return as a slice of points.
// https://www.redblobgames.com/pathfinding/a-star/introduction.html was an excellent guide for this.
func (r *Grid) Path(start image.Point, goal image.Point) []image.Point {
	frontier := make(PriorityQueue, 0)
	heap.Push(&frontier, &Item{value: start, priority: 0})
	cameFrom := map[image.Point]image.Point{
		start: {},
	}
	costSoFar := map[image.Point]int{
		start: 0,
	}
	for len(frontier) > 0 {
		current := heap.Pop(&frontier).(*Item).value
		if current == goal {
			break
		}
		for _, next := range r.Neighbors(current) {
			newCost := costSoFar[current] + int(r.Get(next.X, next.Y))
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				distance := helpers.AbsInt(goal.X-next.X) + helpers.AbsInt(goal.Y-next.Y)
				priority := newCost + distance
				heap.Push(&frontier, &Item{value: next, priority: priority})
				cameFrom[next] = current
			}
		}
	}
	var path []image.Point
	current := goal
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	return path
}

func (r *Grid) PathCost(path []image.Point) int {
	cost := 0
	for _, p := range path {
		cost += int(r.Get(p.X, p.Y))
	}
	return cost
}

func (r *Grid) Neighbors(p image.Point) []image.Point {
	n := make([]image.Point, 0, 4)
	if p.X > 0 {
		n = append(n, image.Point{X: p.X - 1, Y: p.Y})
	}
	if p.X < r.Size()-1 {
		n = append(n, image.Point{X: p.X + 1, Y: p.Y})
	}
	if p.Y > 0 {
		n = append(n, image.Point{X: p.X, Y: p.Y - 1})
	}
	if p.Y < r.Size()-1 {
		n = append(n, image.Point{X: p.X, Y: p.Y + 1})
	}
	return n
}

// SuperSize returns a new grid that is 5x larger in both dimensions, as per part 2 of the puzzle.
func (r *Grid) SuperSize() *Grid {
	getValue := func(x, y int) uint8 {
		distance := y/r.Size() + x/r.Size()
		baseX := x % r.Size()
		baseY := y % r.Size()
		value := r.Get(baseX, baseY) + uint8(distance)
		if value > 9 {
			// 10 -> 1, 11 -> 2, ...
			value = value%10 + 1
		}
		return value
	}
	size := r.Size() * 5
	grid := NewGrid(size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			grid.Set(x, y, getValue(x, y))
		}
	}
	return grid
}
