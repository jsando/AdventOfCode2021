package day17

import (
	"reflect"
	"testing"
)

func TestComputeStep(t *testing.T) {
	tests := []struct {
		input Probe
		want  Probe
	}{
		{
			input: Probe{0, 0, 7, 2},
			want:  Probe{7, 2, 6, 1},
		},
		{
			input: Probe{0, 0, 0, 0},
			want:  Probe{0, 0, 0, -1},
		},
		{
			input: Probe{0, 0, -1, 1},
			want:  Probe{-1, 1, 0, 0},
		},
	}
	for _, test := range tests {
		got := computeStep(test.input)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("want: %v, got %v", test.want, got)
		}
	}
}

func TestParseTarget(t *testing.T) {
	x1, y1, x2, y2 := parseTarget("target area: x=20..30, y=-10..-5")
	if x1 != 20 || x2 != 30 || y1 != -10 || y2 != -5 {
		t.Errorf("expected 20,30 -10,-5 got %d,%d %d,%d", x1, x2, y1, y2)
	}
}

func TestInTarget(t *testing.T) {
	tests := []struct {
		x1, y1, x2, y2 int
		dx, dy         int
		want           bool
	}{
		{20, -10, 30, -5, 7, 2, true},
		{20, -10, 30, -5, 6, 3, true},
		{20, -10, 30, -5, 9, 0, true},
		{20, -10, 30, -5, 17, -4, false},
	}
	for _, test := range tests {
		got := inRange(test.x1, test.y1, test.x2, test.y2, test.dx, test.dy)
		if got != test.want {
			t.Errorf("for %v wanted: %v, got %v", test, test.want, got)
		}
	}
}

//func TestPart1(t *testing.T) {
//	maxY := part1("example-input.txt")
//	if maxY != 45 {
//		t.Errorf("expected 45, got %d", maxY)
//	}
//}

func TestFindMinMax(t *testing.T) {
	got := findMinSum(20)
	if got != 6 {
		t.Errorf("expected 6, got %d", got)
	}
	got = findMinSum(155)
	if got != 18 {
		t.Errorf("expected 6, got %d", got)
	}
}
