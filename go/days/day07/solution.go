package day07

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func getCrabsPerPosition(positions []int) []int {
	maxX := utils.MaxSlice(positions)
	crabsPerPositionArray := make([]int, maxX+1)
	for _, position := range positions {
		crabsPerPositionArray[position] += 1
	}
	return crabsPerPositionArray
}

func getNumMoves(crabsPerPosition []int, targetPosition int) int {
	totalMoves := 0
	for position, numCrabs := range crabsPerPosition {
		numMoves := numCrabs * utils.IntAbs(targetPosition-position)
		totalMoves += numMoves
	}
	return totalMoves
}

func doPart1(crabsPerPosition []int) int {
	// bestPosition := 0
	minMoves := -1
	for targetPosition := range crabsPerPosition {
		numMoves := getNumMoves(crabsPerPosition, targetPosition)
		fmt.Printf("target position %v num moves %v\n", targetPosition, numMoves)
		if numMoves < minMoves || minMoves < 0 {
			// bestPosition = targetPosition
			minMoves = numMoves
		}
	}
	return minMoves
}

func doPart2() int {
	return 0
}

func Run(path string) {
	initialPositions := utils.ParseStringAsIntList(utils.ReadFileAsString(path), ",")
	crabsPerPosition := getCrabsPerPosition(initialPositions)
	fmt.Printf("input: %v\n", initialPositions)
	answer1 := doPart1(crabsPerPosition)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
