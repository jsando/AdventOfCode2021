package day04

import "testing"

func TestLoadGame(t *testing.T) {
	game := NewGameFromFile("example-input.txt")
	if len(game.Boards) != 3 {
		t.Errorf("expected 3 boards, got %d", len(game.Boards))
	}
}

func TestCallNumbers(t *testing.T) {
	game := NewGameFromFile("example-input.txt")
	for i := 0; i < 5; i++ {
		game.CallNext()
	}
	if game.Boards[0].IsWinner() {
		t.Error("you are not a winner")
	}
	called := [][]int{
		{0, 3},
		{1, 3},
		{2, 1},
		{2, 4},
		{3, 4},
	}
	for _, coord := range called {
		if !game.Boards[0].Called[coord[0]][coord[1]] {
			t.Errorf("called number not marked")
		}
	}
	for i := 0; i < 6; i++ {
		game.CallNext()
	}
	if game.Boards[0].IsWinner() {
		t.Error("you are not a winner")
	}
	game.CallNext()
	if game.Boards[0].IsWinner() {
		t.Error("you are not a winner")
	}
	if game.Boards[1].IsWinner() {
		t.Error("you are not a winner")
	}
	if !game.Boards[2].IsWinner() {
		t.Error("but you ARE a winner!")
	}
	if game.Boards[2].Score() != 4512 {
		t.Errorf("expected 4512, got %d", game.Boards[2].Score())
	}
}

func Test_Part1(t *testing.T) {
	score := part1("example-input.txt")
	if score != 4512 {
		t.Errorf("expected 4512, got %d", score)
	}
}

func Test_Part2(t *testing.T) {
	score := part2("example-input.txt")
	if score != 1924 {
		t.Errorf("expected 1924, got %d", score)
	}
}
