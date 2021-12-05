package day03

import (
	"testing"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func TestFindMostCommandBit(t *testing.T) {
	inputs := helpers.FileToStringSlice("test1.txt")
	ones := []string{"1", "0", "1", "1", "0"}
	for bit := 0; bit < len(inputs[0]); bit++ {
		mcb := findMostCommonBit(inputs, bit)
		if mcb != ones[bit] {
			t.Errorf("wrong")
		}
	}
}

func TestFindGamma(t *testing.T) {
	v := findGamma([]string{"100", "101", "011"})
	if v != "101" {
		t.Errorf("got %s, expected 101", v)
	}
}

func TestInvertBinaryString(t *testing.T) {
	if invertBinaryString("10110") != "01001" {
		t.Error()
	}
}

func TestPart1Example(t *testing.T) {
	report := helpers.FileToStringSlice("test1.txt")
	power := part1(report)
	if power != 198 {
		t.Errorf("Expected 198, got %d", power)
	}
}

func TestFindOxygenRating(t *testing.T) {
	inputs := helpers.FileToStringSlice("test1.txt")
	oxygen := findOxygenRating(inputs)
	if oxygen != "10111" {
		t.Errorf("expected 10111, got %s", oxygen)
	}
}

func TestFindC02Rating(t *testing.T) {
	inputs := helpers.FileToStringSlice("test1.txt")
	oxygen := findC02Rating(inputs)
	if oxygen != "01010" {
		t.Errorf("expected 01010, got %s", oxygen)
	}
}

func TestPart2Example(t *testing.T) {
	inputs := helpers.FileToStringSlice("test1.txt")
	rating := part2(inputs)
	if rating != 230 {
		t.Errorf("Expected 230, got %d", rating)
	}
}
