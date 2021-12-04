package day04

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

type bingoGrid struct {
	values [][]int
	status [][]bool
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

func doPart1(draw []int, grids []bingoGrid) int {
	var winningScore, winningNumber int
Loop:
	for _, number := range draw {
		for _, grid := range grids {
			grid.applyNumber(number)
			if grid.isWinning() {
				winningScore = grid.getScore()
				winningNumber = number
				break Loop
			}
		}
	}
	return winningNumber * winningScore
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	draw := utils.ParseStringAsIntList(input[0], ",")
	var grids []bingoGrid
	for _, gridBlob := range input[1:] {
		grids = append(grids, createGrid(gridBlob))
	}

	answer1 := doPart1(draw, grids)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
