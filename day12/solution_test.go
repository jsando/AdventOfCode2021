package day12

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadMap(t *testing.T) {
	caves := NewCaveMapFromFile("example-input.txt")
	fmt.Printf("%v\n", caves)
	connections := caves.GetConnections("A")
	if !reflect.DeepEqual(connections, []string{"c", "b", "end"}) {
		t.Errorf("expected A to connection to start,b,end got '%s'", connections)
	}
}

func TestIsSmallCave(t *testing.T) {
	tests := []struct {
		name  string
		small bool
	}{
		{"a", true},
		{"A", false},
		{"xy", true},
		{"xY", false},
		{"YY", false},
	}
	for _, test := range tests {
		if smallCave(test.name) != test.small {
			t.Errorf("'%s' failed small cave test", test.name)
		}
	}
}

func TestPaths(t *testing.T) {
	tests := []struct {
		caves []string
		from  string
		paths []string
	}{
		{
			caves: []string{"start-end"},
			from:  "start",
			paths: []string{"start,end"},
		},
		{
			caves: []string{"start-end"},
			from:  "end",
			paths: []string{},
		},
		{
			caves: []string{"start-a", "a-end"},
			from:  "a",
			paths: []string{"a,end"},
		},
		{
			caves: []string{"start-a", "a-end"},
			from:  "start",
			paths: []string{"start,a,end"},
		},
		{
			caves: []string{"start-a", "a-b", "b-end"},
			from:  "start",
			paths: []string{"start,a,b,end"},
		},
		{
			caves: []string{"start-a", "a-b", "a-c", "c-b", "b-end"},
			from:  "start",
			paths: []string{"start,a,b,end", "start,a,c,b,end"},
		},
	}
	for _, test := range tests {
		caves := NewMapFromSlice(test.caves)
		paths := caves.PathsFrom(map[string]bool{}, test.from)
		if !reflect.DeepEqual(paths, test.paths) {
			t.Errorf("expected %v, got %v", test.paths, paths)
		}
	}
}

func TestPart1(t *testing.T) {
	paths := part1("example-input.txt")
	if len(paths) != 10 {
		t.Errorf("Expected 10 paths, got %d", len(paths))
	}
}

func TestPart2(t *testing.T) {
	paths := part2("example-input.txt")
	if len(paths) != 36 {
		t.Errorf("Expected 36 paths, got %d", len(paths))
	}
}
