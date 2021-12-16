package day11

import "testing"

func TestLoadInput(t *testing.T) {
	cavern := NewOctopusCavernFromFile("example-input.txt")
	tests := []struct {
		x, y, energy int
	}{
		{x: 0, y: 0, energy: 5},
		{x: 9, y: 1, energy: 1},
		{x: 4, y: 3, energy: 3},
		{x: 9, y: 9, energy: 6},
	}
	for _, test := range tests {
		gotEnergy := cavern.EnergyAt(test.x, test.y)
		if gotEnergy != test.energy {
			t.Errorf("at %d,%d expected %d got %d", test.x, test.y, test.energy, gotEnergy)
		}
	}
}

func TestStep(t *testing.T) {
	cavern := NewOctopusCavernFromFile("example-input.txt")
	cavern.Step()
	flashes, _ := cavern.Step() // the sample input takes 2 steps to flash
	tests := []struct {
		x, y, energy int
	}{
		{x: 2, y: 0, energy: 0},
		{x: 1, y: 1, energy: 0},
		{x: 1, y: 8, energy: 0},
		{x: 3, y: 9, energy: 0},
	}
	for _, test := range tests {
		gotEnergy := cavern.EnergyAt(test.x, test.y)
		if gotEnergy != test.energy {
			t.Errorf("at %d,%d expected %d got %d", test.x, test.y, test.energy, gotEnergy)
		}
	}
	if flashes != 35 {
		t.Errorf("expected 35 flashes total but got %d", flashes)
	}
}

func TestPart1(t *testing.T) {
	flashes := part1("example-input.txt")
	if flashes != 1656 {
		t.Errorf("expected 1656, got %d", flashes)
	}
}

func TestAllFlashed(t *testing.T) {
	cavern := NewOctopusCavernFromFile("example-input2.txt")
	_, allFlashed := cavern.Step()
	if !allFlashed {
		t.Errorf("expected all to flash")
	}
}

func TestPart2(t *testing.T) {
	allFlashStep := part2("example-input.txt")
	if allFlashStep != 195 {
		t.Errorf("expected 195, got %d", allFlashStep)
	}
}
