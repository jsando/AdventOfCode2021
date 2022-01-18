package day13

import (
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
	"strings"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: \n%s\n", part2(inputpath))
}

func part1(inputpath string) int {
	paper := NewPaperFromFile(inputpath)
	paper.Fold(paper.folds[0])
	return paper.PointsVisible()
}

func part2(inputpath string) string {
	paper := NewPaperFromFile(inputpath)
	for _, fold := range paper.folds {
		paper.Fold(fold)
	}
	return paper.String()
}

func foldXY(x, y, foldX, foldY int) (newX, newY int) {
	newX = x
	newY = y
	if foldY != 0 && y > foldY {
		newY = 2*foldY - y
	}
	if foldX != 0 && x >= foldX {
		newX = 2*foldX - x
	}
	return
}

type Point struct {
	x, y int
}

const PaperSize = 2048

type Paper struct {
	points [PaperSize * PaperSize]bool
	folds  []Point
}

func NewPaperFromFile(inputpath string) *Paper {
	paper := &Paper{folds: []Point{}}
	lines := helpers.FileToStringSlice(inputpath)
	addPoints := true
	for _, line := range lines {
		if len(line) == 0 {
			addPoints = false
			continue
		}
		if addPoints {
			var x, y int
			conv, err := fmt.Sscanf(line, "%d,%d", &x, &y)
			if err != nil {
				panic(err)
			}
			if conv != 2 {
				panic("didn't scan 2 ints")
			}
			paper.Set(x, y, true)
		} else {
			var axis rune
			var dist int
			conv, err := fmt.Sscanf(line, "fold along %c=%d", &axis, &dist)
			if err != nil {
				panic(err)
			}
			if conv != 2 {
				panic("didn't scan 2")
			}
			fold := Point{}
			if axis == 'x' {
				fold.x = dist
			} else {
				fold.y = dist
			}
			paper.folds = append(paper.folds, fold)
		}
	}
	return paper
}

func (p *Paper) PointsVisible() int {
	count := 0
	for _, p := range p.points {
		if p {
			count++
		}
	}
	return count
}

func (p *Paper) Set(x int, y int, set bool) {
	p.points[y*PaperSize+x] = set
}

func (p *Paper) Get(x int, y int) bool {
	return p.points[y*PaperSize+x]
}

func (p *Paper) FoldCount() int {
	return len(p.folds)
}

func (p *Paper) Fold(fold Point) {
	for y := 0; y < PaperSize; y++ {
		for x := 0; x < PaperSize; x++ {
			if p.Get(x, y) {
				nx, ny := foldXY(x, y, fold.x, fold.y)
				if nx != x || ny != y {
					p.Set(nx, ny, true)
					p.Set(x, y, false)
				}
			}
		}
	}
}

func (p *Paper) String() string {
	var sb strings.Builder
	blankLineCount := 0
	for y := 0; y < PaperSize; y++ {
		spaceCount := 0
		wroteData := false
		for x := 0; x < PaperSize; x++ {
			if p.Get(x, y) {
				for line := 0; line < blankLineCount; line++ {
					sb.WriteByte('\n')
				}
				blankLineCount = 0
				for space := 0; space < spaceCount; space++ {
					sb.WriteRune(' ')
				}
				spaceCount = 0
				sb.WriteRune('#')
				wroteData = true
			} else {
				spaceCount++
			}
		}
		if wroteData {
			sb.WriteByte('\n')
		} else {
			blankLineCount++
		}
	}
	return sb.String()
}
