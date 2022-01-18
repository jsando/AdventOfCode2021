package helpers

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func TextToIntSlice(text string) []int {
	ints := []int{}
	for _, line := range strings.Split(strings.TrimSpace(text), "\n") {
		if len(line) != 0 {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ints = append(ints, val)
		}
	}
	return ints
}

func FileToIntSlice(path string) []int {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return TextToIntSlice(string(bytes))
}

func TextToStringSlice(text string) []string {
	return strings.Split(strings.TrimSpace(text), "\n")
}

func FileToStringSlice(path string) []string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return TextToStringSlice(string(bytes))
}

func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func StringToIntSlice(text string, sep string) []int {
	values := make([]int, 0)
	text = strings.TrimSpace(text)
	for _, s := range strings.Split(text, sep) {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	return values
}

func MinMaxIntSlice(s []int) (min, max int) {
	if len(s) == 0 {
		panic("empty slice!")
	}
	for i, v := range s {
		if i == 0 || v < min {
			min = v
		}
		if i == 0 || v > max {
			max = v
		}
	}
	return
}
