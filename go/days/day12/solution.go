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
	neighbours []string
}

func (cave *caveNode) addNeighbour(name string) {
	cave.neighbours = append(cave.neighbours, name)
}

func createCave(name string) *caveNode {
	asRunes := []rune(name)
	neighbours := make([]string, 0, 50)
	return &caveNode{
		name:       name,
		isSmall:    unicode.IsLower(asRunes[0]),
		neighbours: neighbours,
	}
}

func getOrCreate(name string, caves map[string]*caveNode) *caveNode {
	cave, caveExists := caves[name]
	if !caveExists {
		cave = createCave(name)
		caves[name] = cave
	}
	return cave
}

func buildCaves(segments []string) map[string]*caveNode {
	caves := make(map[string]*caveNode)
	for _, segment := range segments {
		names := strings.Split(segment, "-")
		name1, name2 := names[0], names[1]
		cave1 := getOrCreate(name1, caves)
		cave2 := getOrCreate(name2, caves)
		cave1.addNeighbour(cave2.name)
		cave2.addNeighbour(cave1.name)
	}
	return caves
}

func copyAndAdd(path []string, caveName string) []string {
	newPath := make([]string, 0, len(path)+1)
	copy(newPath, path)
	return append(newPath, caveName)
}

func traversePart1(caveName string, caves map[string]*caveNode, visited []string, path []string, endPaths *[][]string) {
	// Bail if we reach a cave twice
	if utils.StringInSlice(caveName, visited) {
		return
	}

	newPath := copyAndAdd(path, caveName)

	// Bail and add a new end path if we reach the end
	if caveName == "end" {
		*endPaths = append(*endPaths, newPath)
		return
	}
	cave := caves[caveName]
	if cave.isSmall {
		visited = append(visited, caveName)
	}
	for _, neighbourName := range cave.neighbours {
		traversePart1(neighbourName, caves, visited, newPath, endPaths)
	}

}

func traversePart2(caveName string, caves map[string]*caveNode, visited []string, nameOfCaveVisitedTwice string, path []string, endPaths *[][]string, isRoot bool) {
	caveAlreadyVisited := utils.StringInSlice(caveName, visited)

	if (caveName == "start" && !isRoot) || // start
		caveName == nameOfCaveVisitedTwice || // reached 3 times
		nameOfCaveVisitedTwice != "" && caveAlreadyVisited { // reached two caves twice
		return
	}

	newPath := copyAndAdd(path, caveName)

	// Bail and add a new end path if we reach the end
	if caveName == "end" {
		*endPaths = append(*endPaths, newPath)
		return
	}

	cave := caves[caveName]
	if cave.isSmall {
		if caveAlreadyVisited {
			nameOfCaveVisitedTwice = caveName
		} else {
			visited = append(visited, caveName)
		}
	}
	for _, neighbourName := range cave.neighbours {
		traversePart2(neighbourName, caves, visited, nameOfCaveVisitedTwice, newPath, endPaths, false)
	}

}

func doPart1(caves map[string]*caveNode) int {
	endPaths := make([][]string, 0, 200000)
	traversePart1("start", caves, make([]string, 0), make([]string, 0), &endPaths)
	return len(endPaths)
}

func doPart2(caves map[string]*caveNode) int {
	endPaths := make([][]string, 0, 200000)
	traversePart2("start", caves, make([]string, 0), "", make([]string, 0), &endPaths, true)
	return len(endPaths)
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	caves := buildCaves(input)
	answer1 := doPart1(caves)
	answer2 := doPart2(caves)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
