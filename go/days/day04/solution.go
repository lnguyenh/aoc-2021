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

func doPart2(draw []int, grids []bingoGrid) int {
	var results []int
	for _, number := range draw {
		for _, grid := range grids {
			if grid.isWinning() {
				continue
			}
			grid.applyNumber(number)
			if grid.isWinning() {
				results = append(results, grid.getScore()*number)
			}
		}
	}

	return results[len(results)-1]
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	draw := utils.ParseStringAsIntList(input[0], ",")
	var gridsPart1, gridsPart2 []bingoGrid
	for _, gridBlob := range input[1:] {
		gridsPart1 = append(gridsPart1, createGrid(gridBlob))
		gridsPart2 = append(gridsPart2, createGrid(gridBlob))
	}

	answer1 := doPart1(draw, gridsPart1)
	answer2 := doPart2(draw, gridsPart2)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
