package day12

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
	"unicode"
)

type caveNode struct {
	name       string
	isSmall    bool
	neighbours map[string]bool
}

func (cave *caveNode) add(neighbourCave caveNode) {
	cave.neighbours[neighbourCave.name] = true
}

func createCave(name string) caveNode {
	asRunes := []rune(name)
	neighbours := make(map[string]bool)
	return caveNode{
		name:       name,
		isSmall:    unicode.IsLower(asRunes[0]),
		neighbours: neighbours,
	}
}

func traverse(caveName string, caves map[string]caveNode, visited []string, path []string, endPaths *[][]string) {
	newPath := make([]string, len(path))
	copy(newPath, path)
	newPath = append(newPath, caveName)
	if caveName == "end" {
		*endPaths = append(*endPaths, newPath)
		return
	}
	if utils.StringInSlice(caveName, visited) {
		return
	} else {
		cave := caves[caveName]
		if cave.isSmall {
			visited = append(visited, caveName)
		}
		for neighbourName := range cave.neighbours {
			traverse(neighbourName, caves, visited, newPath, endPaths)
		}
	}
}

func doPart1(caves map[string]caveNode) int {
	endPaths := make([][]string, 0)
	traverse("start", caves, make([]string, 0), make([]string, 0), &endPaths)
	for _, path := range endPaths {
		fmt.Printf("%v\n", path)
	}
	return len(endPaths)
}

func doPart2() int {
	return 0
}

func buildCaves(segments []string) map[string]caveNode {
	caves := make(map[string]caveNode)
	for _, segment := range segments {
		names := strings.Split(segment, "-")
		name1, name2 := names[0], names[1]
		_, cave1Exists := caves[name1]
		_, cave2Exists := caves[name2]
		if !cave1Exists {
			newCave := createCave(name1)
			caves[name1] = newCave
		}
		if !cave2Exists {
			newCave := createCave(name2)
			caves[name2] = newCave
		}
		cave1 := caves[name1]
		cave2 := caves[name2]
		cave1.add(cave2)
		cave2.add(cave1)
	}
	return caves
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	caves := buildCaves(input)
	fmt.Printf("input: %v\n", len(caves))
	answer1 := doPart1(caves)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
