package day08

import (
	"fmt"
	"strings"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", part1(inputpath))
	fmt.Printf("Part 2: %d\n", part2(inputpath))
}

func part1(inputpath string) int {
	sum := 0
	lines := helpers.FileToStringSlice(inputpath)
	for _, line := range lines {
		digits := strings.TrimSpace(strings.Split(line, "|")[1])
		sum += countKnown(digits)
	}
	return sum
}

func part2(inputpath string) int {
	sum := 0
	lines := helpers.FileToStringSlice(inputpath)
	for _, line := range lines {
		sum += decodeSegments(line)
	}
	return sum
}

func countKnown(input string) int {
	list := strings.Split(input, " ")
	count := 0
	for _, segments := range list {
		switch len(segments) {
		case 2, 3, 4, 7:
			count++
		}
	}
	return count
}

func mapGroupings(groups string) map[string]int {
	mappings := map[string]int{}

	// 1, 7, 4, and 8 are easily deduced just from the character count
	var one, four string
	for _, group := range strings.Split(groups, " ") {
		sorted := sortSegments(group)
		value := -1
		switch len(sorted) {
		case 2:
			one = group
			value = 1
		case 3:
			value = 7
		case 4:
			four = group
			value = 4
		case 7:
			value = 8

		}
		mappings[sorted] = value
	}

	// Solve the 6 char segments:
	//  9: contains all the chars from 4
	//  0: has all the chars from 1
	//  6: leftover
	for k := range mappings {
		if len(k) == 6 {
			if countCommon(k, four) == 4 {
				mappings[k] = 9
			} else if countCommon(k, one) == 2 {
				mappings[k] = 0
			} else {
				mappings[k] = 6
			}
		}
	}

	// Solve the 5 char segments:
	//  3: has all letters from 1 but have to identify it first so can then just have 2 and 5
	for k := range mappings {
		if len(k) == 5 && countCommon(k, one) == 2 {
			mappings[k] = 3
		}
	}

	//  2: 2 letters in common with 4
	//  5: 3 letters in common with 4
	for k, v := range mappings {
		if len(k) == 5 && v == -1 {
			if countCommon(k, four) == 2 {
				mappings[k] = 2
			} else {
				mappings[k] = 5
			}
		}
	}

	return mappings
}

// Return true if s contains all characters from subset
func countCommon(s1, s2 string) int {
	common := 0
	for _, ch := range s1 {
		for _, ch2 := range s2 {
			if ch == ch2 {
				common++
				break
			}
		}
	}
	return common
}

func sortSegments(s string) string {
	var alphabet [7]bool
	for _, ch := range s {
		alphabet[ch-'a'] = true
	}
	out := ""
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] {
			out += string(rune('a' + i))
		}
	}
	return out
}

func decodeSegments(line string) int {
	value := 0
	parts := strings.Split(line, "|")
	groups := strings.TrimSpace(parts[0])
	mapping := mapGroupings(groups)
	for _, group := range strings.Split(parts[1], " ") {
		group = sortSegments(group)
		digit := mapping[group]
		if digit == -1 {
			panic(fmt.Sprintf("unkown group %s", group))
		}
		value = value*10 + digit
	}
	return value
}
