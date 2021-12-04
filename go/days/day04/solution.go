package day04

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

type grid struct {
	values   [][]int
	statuses [][]bool
}

func createGrid(blob string) grid {
	var lines [][]int
	for _, line := range strings.Split(blob, "\n") {
		lines = append(lines, utils.ParseStringAsIntList(line, " "))
	}
	newGrid := grid{
		values: lines,
		statuses: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}
	return newGrid
}

func doPart1() int {
	return 0
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	draw := utils.ParseStringAsIntList(input[0], ",")
	var grids []grid
	for _, gridBlob := range input[1:] {
		grids = append(grids, createGrid(gridBlob))
	}

	fmt.Println(grids)
	fmt.Println(draw)
	answer1 := doPart1()
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
