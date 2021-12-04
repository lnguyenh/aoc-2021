package day04

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

type bingoGrid struct {
	values    [][]int
	status    [][]bool
	rank      int
	winNumber int
}

func createGrid(blob string) bingoGrid {
	var lines [][]int
	for _, line := range strings.Split(blob, "\n") {
		lines = append(lines, utils.ParseStringAsIntList(line, " "))
	}
	newGrid := bingoGrid{
		values: lines,
		status: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}
	return newGrid
}

func (grid bingoGrid) applyNumber(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid.status[i][j] || grid.values[i][j] != number {
				continue
			}
			grid.status[i][j] = true
		}
	}
}

func (grid bingoGrid) isWinning() bool {
	for i := 0; i < 5; i++ {
		rowContainsFalse := false
		for j := 0; j < 5; j++ {
			if !grid.status[i][j] {
				rowContainsFalse = true
				break
			}
		}
		if !rowContainsFalse {
			return true
		}
	}

	for j := 0; j < 5; j++ {
		columnContainsFalse := false
		for i := 0; i < 5; i++ {
			if !grid.status[i][j] {
				columnContainsFalse = true
				break
			}
		}
		if !columnContainsFalse {
			return true
		}
	}

	return false
}

func (grid bingoGrid) getScore() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !grid.status[i][j] {
				score += grid.values[i][j]
			}
		}
	}
	return score
}

func createGrids(gridBlobs []string) []bingoGrid {
	var grids []bingoGrid
	for _, gridBlob := range gridBlobs {
		grids = append(grids, createGrid(gridBlob))
	}
	return grids
}

func runGame(draw []int, grids []bingoGrid) {
	rankCounter := 1
	numGrids := len(grids)
	for _, number := range draw {
		for i := 0; i < numGrids; i++ {
			grid := &grids[i]
			if grid.rank > 0 {
				continue
			}
			grid.applyNumber(number)
			if grid.isWinning() {
				grid.rank = rankCounter
				grid.winNumber = number
				rankCounter++
			}
		}
	}
}

func doPart1(grids []bingoGrid) int {
	for _, grid := range grids {
		if grid.rank == 1 {
			return grid.getScore() * grid.winNumber
		}
	}
	return 0
}

func doPart2(grids []bingoGrid) int {
	lastRank := len(grids)
	for _, grid := range grids {
		if grid.rank == lastRank {
			return grid.getScore() * grid.winNumber
		}
	}
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")

	draw := utils.ParseStringAsIntList(input[0], ",")
	grids := createGrids(input[1:])

	runGame(draw, grids)

	answer1 := doPart1(grids)
	answer2 := doPart2(grids)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
