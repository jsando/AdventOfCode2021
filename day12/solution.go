package day12

import (
	"fmt"
	"strings"

	"github.com/jsando/AdventOfCode2021/helpers"
)

func Run(inputpath string) {
	paths := part1(inputpath)
	fmt.Printf("Part 1: %d\n", len(paths))
	paths = part2(inputpath)
	fmt.Printf("Part 2: %d\n", len(paths))
}

func part1(inputpath string) []string {
	caves := NewCaveMapFromFile(inputpath)
	return caves.PathsFrom(map[string]bool{}, "start")
}

func part2(inputpath string) []string {
	caves := NewCaveMapFromFile(inputpath)
	return caves.Paths2()
}

type CaveMap struct {
	caves map[string][]string
}

func NewCaveMapFromFile(inputpath string) *CaveMap {
	return NewMapFromSlice(helpers.FileToStringSlice(inputpath))
}

func NewMapFromSlice(lines []string) *CaveMap {
	caves := &CaveMap{
		caves: make(map[string][]string),
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		names := strings.Split(line, "-")
		fromCave := names[0]
		toCave := names[1]
		caves.Connect(fromCave, toCave)
		caves.Connect(toCave, fromCave)
	}
	return caves
}

func (c *CaveMap) GetConnections(fromCave string) []string {
	return c.caves[fromCave]
}

func (c *CaveMap) Connect(fromCave, toCave string) {
	if fromCave == "end" || toCave == "start" {
		return
	}
	connections := c.caves[fromCave]
	if connections == nil {
		c.caves[fromCave] = []string{}
	}
	c.caves[fromCave] = append(c.caves[fromCave], toCave)
}

func smallCave(s string) bool {
	return strings.ToLower(s) == s
}

func (c *CaveMap) PathsFrom(visited map[string]bool, start string) []string {
	if start == "end" {
		return []string{}
	}
	choices := c.GetConnections(start)
	paths := []string{}
	if smallCave(start) {
		visited[start] = true
	}
	for _, choice := range choices {
		if choice == "end" {
			paths = append(paths, start+",end")
		} else if !visited[choice] {
			visitCopy := map[string]bool{}
			for k, v := range visited {
				visitCopy[k] = v
			}
			subpaths := c.PathsFrom(visitCopy, choice)
			for _, sp := range subpaths {
				if strings.HasSuffix(sp, ",end") {
					paths = append(paths, start+","+sp)
				}
			}
		}
	}
	return paths
}

func (c *CaveMap) Paths2() []string {
	paths := map[string]bool{}
	smallCaves := []string{}
	for caveName := range c.caves {
		if smallCave(caveName) {
			smallCaves = append(smallCaves, caveName)
		}
	}
	for _, visitThisOneTwice := range smallCaves {
		allowed := map[string]int{}
		for caveName := range c.caves {
			count := -1
			if caveName == visitThisOneTwice {
				count = 2
			} else if smallCave(caveName) {
				count = 1
			}
			allowed[caveName] = count
		}
		foundPaths := c.PathsFrom2(allowed, "start")
		for _, path := range foundPaths {
			paths[path] = true
		}
	}
	pathSlice := []string{}
	for path := range paths {
		pathSlice = append(pathSlice, path)
	}
	return pathSlice
}

func (c *CaveMap) PathsFrom2(allowed map[string]int, start string) []string {
	if start == "end" {
		return []string{}
	}
	choices := c.GetConnections(start)
	paths := []string{}
	allowed[start]--
	for _, choice := range choices {
		if choice == "end" {
			paths = append(paths, start+",end")
		} else if allowed[choice] != 0 {
			visitCopy := map[string]int{}
			for k, v := range allowed {
				visitCopy[k] = v
			}
			subpaths := c.PathsFrom2(visitCopy, choice)
			for _, sp := range subpaths {
				if strings.HasSuffix(sp, ",end") {
					paths = append(paths, start+","+sp)
				}
			}
		}
	}
	return paths
}
