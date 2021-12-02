package day02

import (
	"fmt"
	"testing"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func TestHorizontalMovement1(t *testing.T) {
	var sub Submarine
	if sub.Horizontal != 0 {
		t.Error()
	}
	sub.Move("forward 5")
	if sub.Horizontal != 5 {
		t.Error()
	}
	sub.Move("forward 5")
	if sub.Horizontal != 10 {
		t.Error()
	}
}

func TestVerticalMovement1(t *testing.T) {
	var sub Submarine
	if sub.Depth != 0 {
		t.Error()
	}
	sub.Move("down 5")
	if sub.Depth != 5 {
		t.Error()
	}
	sub.Move("down 5")
	if sub.Depth != 10 {
		t.Error()
	}
	sub.Move("up 7")
	if sub.Depth != 3 {
		t.Error()
	}
}

func TestPart1(t *testing.T) {
	// Provided as the example in the README
	testInput := `forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2
	`
	var sub Submarine
	sub.MoveSeries(helpers.TextToStringSlice(testInput))
	if sub.Horizontal != 15 {
		t.Error(fmt.Sprintf("wrong horizontal position: %d", sub.Horizontal))
	}
	if sub.Depth != 10 {
		t.Error(fmt.Sprintf("wrong depth: %d", sub.Depth))
	}
}

func TestVerticalMovement2(t *testing.T) {
	var sub Submarine2
	if sub.Aim != 0 {
		t.Error()
	}
	sub.Move("down 5")
	if sub.Aim != 5 {
		t.Error()
	}
	sub.Move("down 5")
	if sub.Aim != 10 {
		t.Error()
	}
	sub.Move("up 7")
	if sub.Aim != 3 {
		t.Error()
	}
}

func TestHorizontalMovement2(t *testing.T) {
	var sub Submarine2
	if sub.Horizontal != 0 {
		t.Error()
	}
	if sub.Depth != 0 {
		t.Error()
	}
	sub.Move("forward 5")
	if sub.Horizontal != 5 {
		t.Error()
	}
	sub.Move("forward 5")
	if sub.Horizontal != 10 {
		t.Error()
	}
}

func TestDepthChangesWithAim(t *testing.T) {
	var sub Submarine2
	if sub.Depth != 0 {
		t.Error()
	}
	sub.Move("forward 5")
	if sub.Horizontal != 5 || sub.Depth != 0 {
		t.Errorf("%v", sub)
	}
	sub.Move("down 5")
	if sub.Aim != 5 {
		t.Errorf("%v", sub)
	}
	sub.Move("forward 8")
	if sub.Horizontal != 13 || sub.Depth != 40 {
		t.Errorf("%#v", sub)
	}
}
