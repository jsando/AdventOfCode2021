package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	inputs := helpers.FileToStringSlice(inputpath)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	gammaString := findGamma(inputs)
	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(invertBinaryString(gammaString), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(gamma * epsilon)
}

func part2(inputs []string) int {
	oxygenString := findOxygenRating(inputs)
	oxygen, err := strconv.ParseInt(oxygenString, 2, 64)
	if err != nil {
		panic(err)
	}
	c02String := findC02Rating(inputs)
	c02, err := strconv.ParseInt(c02String, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(oxygen * c02)
}

func findGamma(inputs []string) string {
	result := ""
	bitCount := len(inputs[0])
	for bit := 0; bit < bitCount; bit++ {
		result += findMostCommonBit(inputs, bit)
	}
	return result
}

func findMostCommonBit(inputs []string, bit int) string {
	zeroes := 0
	ones := 0
	for _, v := range inputs {
		if v[bit] == '1' {
			ones++
		} else {
			zeroes++
		}
	}
	if ones >= zeroes {
		return "1"
	}
	return "0"
}

func invertBinaryString(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '0':
			return '1'
		case '1':
			return '0'
		}
		return r
	}, s)
}

func findC02Rating(inputs []string) string {
	mask := ""
	for bit := 0; bit < len(inputs[0]); bit++ {
		mcb := findMostCommonBit(inputs, bit)
		lcb := invertBinaryString(mcb)
		mask += lcb
		inputs = keepIfStartsWith(inputs, mask)
		if len(inputs) == 1 {
			break
		}
	}
	if len(inputs) != 1 {
		panic("foo")
	}
	return inputs[0]
}

func findOxygenRating(inputs []string) string {
	mask := ""
	for bit := 0; bit < len(inputs[0]); bit++ {
		mcb := findMostCommonBit(inputs, bit)
		mask += mcb
		inputs = keepIfStartsWith(inputs, mask)
		if len(inputs) == 1 {
			break
		}
	}
	if len(inputs) != 1 {
		panic("foo")
	}
	return inputs[0]
}

func keepIfStartsWith(inputs []string, mask string) []string {
	outputs := make([]string, 0, len(inputs))
	for _, v := range inputs {
		if strings.HasPrefix(v, mask) {
			outputs = append(outputs, v)
		}
	}
	return outputs
}
