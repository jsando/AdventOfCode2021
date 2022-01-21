package day15

import (
	"container/heap"
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
	"image"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath)) // 578 too low, added only right/down constraint and now 583 is correct
	fmt.Printf("Part 2: %d\n", part2(inputpath)) // 2934 too high (WTF!), re-enabled up/left and now get 2927
}

func part1(inputpath string) int {
	grid := NewGridFromFile(inputpath)
	path := grid.Path(image.Point{X: 0, Y: 0}, image.Point{X: grid.Width() - 1, Y: grid.Height() - 1})
	//grid.PrintPath(path)
	return grid.PathCost(path)
}

func part2(inputpath string) int {
	grid := NewGridFromFile(inputpath)
	grid.SuperSize()
	path := grid.Path(image.Point{X: 0, Y: 0}, image.Point{X: grid.Width() - 1, Y: grid.Height() - 1})
	//grid.PrintPath(path)
	return grid.PathCost(path)
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
	frontier := make(PriorityQueue, 0)
	frontier.PushPoint(start, 0)
	cameFrom := map[image.Point]image.Point{
		start: {},
	}
	costSoFar := map[image.Point]int{
		start: 0,
	}
	for len(frontier) > 0 {
		current := frontier.PopPoint()
		if current == goal {
			break
		}
		for _, next := range r.Neighbors(current) {
			newCost := costSoFar[current] + r.CostOf(next)
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := newCost + heuristic(goal, next)
				frontier.PushPoint(next, priority)
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

func (r *RiskMap) PathCost(path []image.Point) int {
	cost := -1
	for _, p := range path {
		cost += r.CostOf(p)
	}
	return cost
}

func heuristic(p1, p2 image.Point) int {
	return helpers.AbsInt(p1.X-p2.X) + helpers.AbsInt(p1.Y-p2.Y)
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
			if ch == 0 {
				panic("just kidding")
			}
			rowCopy = append(rowCopy, ' ')
		}
		grid = append(grid, rowCopy)
	}
	for _, p := range path {
		grid[p.Y][p.X] = r.grid[p.Y][p.X]
	}
	fmt.Printf("\nFound path:\n")
	for _, row := range grid {
		for _, ch := range row {
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
}

func (r *RiskMap) CostOf(point image.Point) int {
	return int(r.grid[point.Y][point.X] - '0')
}

// SuperSize expands the map by 5x in both directions as per part2 of the puzzle.
func (r *RiskMap) SuperSize() {
	getValue := func(x, y int) rune {
		distance := y/r.Height() + x/r.Width()
		baseX := x % r.Width()
		baseY := y % r.Height()
		value := int(r.grid[baseY][baseX]-'0') + distance
		if value > 9 { // 10 -> 1, 11 -> 2, 12 -> 3
			value = value%10 + 1
		}
		return rune(value + '0')
	}
	height := r.Height() * 5
	width := r.Width() * 5
	grid := make([][]rune, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			grid[y][x] = getValue(x, y)
		}
	}
	r.grid = grid
}

// Copy-paste from Go standard library example for 'heap'.

// An Item is something we manage in a priority queue.
type Item struct {
	value    image.Point // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	index    int         // The index of the item in the heap (maintained by the heap.Interface methods)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PushPoint(p image.Point, cost int) {
	heap.Push(pq, &Item{value: p, priority: cost})
}

func (pq *PriorityQueue) PopPoint() image.Point {
	item := heap.Pop(pq).(*Item)
	return item.value
}
