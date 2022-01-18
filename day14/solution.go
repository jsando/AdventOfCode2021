package day14

import (
	"fmt"
	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	fmt.Printf("Part 1: %d\n", processExpansion(inputpath, 10))
	fmt.Printf("Part 2: %d\n", processExpansion(inputpath, 40))
}

func processExpansion(inputpath string, steps int) int {
	template, rules := loadRules(inputpath)

	// Initialize count by letter to the initial template, then keep track
	// as letters are inserted in order to compute the final answer.
	letterCount := make([]int, 26)
	for _, ch := range template {
		letterCount[ch-'A']++
	}

	// Keep count by pairs, so never have to store more than nPk (90, given 2P10)
	expansion := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		expansion[pair]++
	}

	for step := 0; step < steps; step++ {
		snapshot := map[string]int{}
		for k, v := range expansion {
			snapshot[k] = v
		}
		for k, v := range snapshot {
			insert := rules[k]
			if len(insert) > 0 {
				expansion[k] -= v
				letterCount[insert[0]-'A'] += v
				k2 := k[0:1] + insert[0:1]
				k3 := insert[0:1] + k[1:2]
				expansion[k2] += v
				expansion[k3] += v
			}
		}
	}

	// Find (max-min)
	var min, max int
	for _, v := range letterCount {
		if v > 0 {
			if min == 0 || v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}
	return max - min
}

func loadRules(inputpath string) (string, map[string]string) {
	lines := helpers.FileToStringSlice(inputpath)
	template := lines[0]
	rules := map[string]string{}
	chars := map[rune]bool{}
	for _, line := range lines[2:] {
		var pair, substitution string
		count, err := fmt.Sscanf(line, "%s -> %s", &pair, &substitution)
		if count != 2 || err != nil {
			panic("scan failed")
		}
		rules[pair] = substitution
		chars[rune(pair[0])] = true
		chars[rune(pair[1])] = true
		chars[rune(substitution[0])] = true
	}
	return template, rules
}
