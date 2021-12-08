package day06

import (
	"reflect"
	"testing"
)

func TestFishTick(t *testing.T) {
	tests := []struct {
		timerIn  int
		timerOut int
		spawn    []int
	}{
		{timerIn: 7, timerOut: 6, spawn: []int{}},
		{timerIn: 6, timerOut: 5, spawn: []int{}},
		{timerIn: 5, timerOut: 4, spawn: []int{}},
		{timerIn: 4, timerOut: 3, spawn: []int{}},
		{timerIn: 3, timerOut: 2, spawn: []int{}},
		{timerIn: 2, timerOut: 1, spawn: []int{}},
		{timerIn: 1, timerOut: 0, spawn: []int{}},
		{timerIn: 0, timerOut: 6, spawn: []int{8}},
	}
	for i, test := range tests {
		out, spawn := fishTick(test.timerIn)
		if out != test.timerOut {
			t.Errorf("in test %d, expected %d got %d", i, test.timerOut, out)
		}
		if !reflect.DeepEqual(test.spawn, spawn) {
			t.Errorf("in test %d, expected %v got %v", i, test.spawn, spawn)
		}
	}
}

func TestSchoolTick(t *testing.T) {
	school := []int{3, 4, 3, 1, 2}
	expected := [][]int{
		{2, 3, 2, 0, 1},
		{1, 2, 1, 6, 0, 8},
		{0, 1, 0, 5, 6, 7, 8},
		{6, 0, 6, 4, 5, 6, 7, 8, 8},
		{5, 6, 5, 3, 4, 5, 6, 7, 7, 8},
	}
	for _, expect := range expected {
		school = schoolTick(school)
		if !reflect.DeepEqual(school, expect) {
			t.Errorf("expected %v, got %v", expect, school)
		}
	}
}

func TestPart1(t *testing.T) {
	count := part1("example-input.txt")
	if count != 5934 {
		t.Errorf("expected 5934, got %d", count)
	}
}

// As often happens, part 2 of the puzzle scales way higher than part 1 and
// forces a more ram/cpu efficient approach.
func TestFishTick2(t *testing.T) {
	//school := []int{3, 4, 3, 1, 2}
	school := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}
	expected := [][]int{
		{1, 1, 2, 1, 0, 0, 0, 0, 0},
		{1, 2, 1, 0, 0, 0, 1, 0, 1},
		{2, 1, 0, 0, 0, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 1, 3, 1, 2},
		{0, 0, 0, 1, 1, 3, 2, 2, 1},
		{0, 0, 1, 1, 3, 2, 2, 1, 0},
	}
	for _, expect := range expected {
		school2Tick(school)
		if !reflect.DeepEqual(school, expect) {
			t.Errorf("expected %v, got %v", expect, school)
		}
	}

}

func TestPart2(t *testing.T) {
	count := part2("example-input.txt")
	if count != 26984457539 {
		t.Errorf("expected 26984457539, got %d", count)
	}
}
