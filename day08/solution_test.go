package day08

import (
	"testing"
)

func TestCountKnown(t *testing.T) {
	tests := []struct {
		input string
		count int
	}{
		{input: "", count: 0},
		{input: "a aa", count: 1},
		{input: "a aa aaa", count: 2},
		{input: "a aa aaa aaaa", count: 3},
		{input: "aaaaaaa", count: 1},
		{input: "a aaaaa aaaaaa", count: 0},
	}
	for _, test := range tests {
		count := countKnown(test.input)
		if count != test.count {
			t.Errorf("for %v, expected %d got %d", test.input, test.count, count)
		}
	}
}

func TestPart1(t *testing.T) {
	count := part1("example-input.txt")
	if count != 26 {
		t.Errorf("expected 26, got %d", count)
	}
}

func TestMapSegments(t *testing.T) {
	numbersByGrouping := mapGroupings("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab")
	tests := []struct {
		grouping string
		value    int
	}{
		{grouping: "ab", value: 1},
		{grouping: "abd", value: 7},
		{grouping: "abef", value: 4},
		{grouping: "abcdefg", value: 8},
		{grouping: "abcdef", value: 9},
		{grouping: "bcdefg", value: 6},
		{grouping: "fbcad", value: 3},
		{grouping: "cdfbe", value: 5},
		{grouping: "gcdfa", value: 2},
	}
	for _, test := range tests {
		got := numbersByGrouping[sortSegments(test.grouping)]
		if got != test.value {
			t.Errorf("for %s expected %d, got %d", test.grouping, test.value, got)
		}
	}
}

func TestSortChars(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "dab", expected: "abd"},
		{in: "cdfgeb", expected: "bcdefg"},
		{in: "acedgfb", expected: "abcdefg"},
	}
	for _, test := range tests {
		got := sortSegments(test.in)
		if got != test.expected {
			t.Errorf("expected %s, got %s", test.expected, got)
		}
	}
}

func TestDecoder(t *testing.T) {
	value := decodeSegments("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	if value != 5353 {
		t.Errorf("expected 5353, got %d", value)
	}
}

func TestPart2(t *testing.T) {
	sum := part2("example-input.txt")
	if sum != 61229 {
		t.Errorf("expected 61229, got %d", sum)
	}
}
