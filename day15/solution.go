package day15

import (
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
	"image"
)

func Run(inputpath string) {

}

type RiskMap struct {
	grid [][]rune
}

func NewGridFromFile(inputpath string) *RiskMap {
	r := &RiskMap{
		grid: make([][]rune, 0),
	}
	lines := helpers.FileToStringSlice(inputpath)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]rune, 0)
		for _, ch := range line {
			row = append(row, ch)
		}
		r.grid = append(r.grid, row)
	}
	return r
}

func (r *RiskMap) Width() int {
	return len(r.grid[0])
}

func (r *RiskMap) Height() int {
	return len(r.grid)
}

func (r *RiskMap) Path(start image.Point, goal image.Point) []image.Point {
	var frontier []image.Point
	frontier = append(frontier, start)
	cameFrom := map[image.Point]image.Point{
		start: {},
	}
	for len(frontier) > 0 {
		current := frontier[0]
		if current == goal {
			break
		}
		frontier = frontier[1:]
		for _, next := range r.Neighbors(current) {
			if _, ok := cameFrom[next]; !ok {
				frontier = append(frontier, next)
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

func (r *RiskMap) Neighbors(p image.Point) []image.Point {
	n := make([]image.Point, 0, 4)
	if p.X > 0 {
		n = append(n, image.Point{X: p.X - 1, Y: p.Y})
	}
	if p.X < r.Width()-1 {
		n = append(n, image.Point{X: p.X + 1, Y: p.Y})
	}
	if p.Y > 0 {
		n = append(n, image.Point{X: p.X, Y: p.Y - 1})
	}
	if p.Y < r.Height()-1 {
		n = append(n, image.Point{X: p.X, Y: p.Y + 1})
	}
	return n
}

func (r *RiskMap) PrintPath(path []image.Point) {
	grid := [][]rune{}
	for _, row := range r.grid {
		rowCopy := []rune{}
		for _, ch := range row {
			rowCopy = append(rowCopy, ch)
		}
		grid = append(grid, rowCopy)
	}
	for _, p := range path {
		grid[p.Y][p.X] = '*'
	}
	fmt.Printf("\nFound path:\n")
	for _, row := range grid {
		for _, ch := range row {
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
}
