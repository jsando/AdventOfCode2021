package day04

import (
	"fmt"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	game := NewGameFromFile(inputpath)
	fmt.Printf("Game has %d boards and %d calls.\n", len(game.Boards), len(game.Called))
	for {
		game.CallNext()
		for _, board := range game.Boards {
			if board.IsWinner() {
				return board.Score()
			}
		}
	}
}

func part2(inputpath string) int {
	game := NewGameFromFile(inputpath)
	var lastBoard *Board
	for {
		game.CallNext()
		allWon := true
		for _, board := range game.Boards {
			if !board.IsWinner() {
				allWon = false
				lastBoard = board
			}
		}
		if allWon {
			break
		}
	}
	if lastBoard == nil {
		panic("lastBoard is nil!")
	}
	return lastBoard.Score()
}

type Board struct {
	Cells      [5][5]int
	Called     [5][5]bool
	LastNumber int
}

type Game struct {
	Called   []int
	NextCall int
	Boards   []*Board
}

func NewGameFromFile(inputpath string) *Game {
	game := &Game{}
	lines := helpers.FileToStringSlice(inputpath)
	game.Called = helpers.StringToIntSlice(lines[0], ",")
	for i := 2; i < len(lines); i += 6 {
		board := &Board{}
		game.Boards = append(game.Boards, board)
		for y := 0; y < 5; y++ {
			s := lines[i+y]
			numbers := helpers.StringToIntSlice(s, " ")
			for x := 0; x < 5; x++ {
				board.Cells[y][x] = numbers[x]
			}
		}
	}
	return game
}

func (g *Game) CallNext() {
	if g.NextCall > len(g.Called) {
		panic("no more calls")
	}
	call := g.Called[g.NextCall]
	g.NextCall++
	for _, board := range g.Boards {
		board.Call(call)
	}
}

func (b *Board) Call(number int) {
	b.LastNumber = number
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.Cells[y][x] == number {
				b.Called[y][x] = true
			}
		}
	}
}

func (b *Board) IsWinner() bool {
	for y := 0; y < 5; y++ {
		rowWinner := true
		for x := 0; x < 5; x++ {
			if !b.Called[y][x] {
				rowWinner = false
				break
			}
		}
		if rowWinner {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		colWinner := true
		for y := 0; y < 5; y++ {
			if !b.Called[y][x] {
				colWinner = false
				break
			}
		}
		if colWinner {
			return true
		}
	}
	return false
}

func (b *Board) Score() int {
	sum := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.Called[y][x] {
				sum += b.Cells[y][x]
			}
		}
	}
	return b.LastNumber * sum
}
