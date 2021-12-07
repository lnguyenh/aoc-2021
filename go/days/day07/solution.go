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

func getFuelPart1(crabsPerPosition []int, targetPosition int) int {
	totalMoves := 0
	for position, numCrabs := range crabsPerPosition {
		numMoves := numCrabs * utils.IntAbs(targetPosition-position)
		totalMoves += numMoves
	}
	return totalMoves
}

func getCostsPart2(crabsPerPosition []int) []int {
	costs := make([]int, len(crabsPerPosition)*utils.MaxSlice(crabsPerPosition))
	accumulator := 0
	for numSteps := range costs {
		costs[numSteps] = numSteps + accumulator
		accumulator += numSteps
	}
	return costs
}

func getFuelPart2(crabsPerPosition []int, targetPosition int, costs []int) int {
	totalFuel := 0
	for position, numCrabs := range crabsPerPosition {
		numMoves := utils.IntAbs(targetPosition - position)
		totalFuel += costs[numMoves] * numCrabs
	}
	return totalFuel
}

func doPart1(crabsPerPosition []int) int {
	minMoves := -1
	for targetPosition := range crabsPerPosition {
		numMoves := getFuelPart1(crabsPerPosition, targetPosition)
		if numMoves < minMoves || minMoves < 0 {
			minMoves = numMoves
		}
	}
	return minMoves
}

func doPart2(crabsPerPosition []int, costs []int) int {
	minFuel := -1
	for targetPosition := range crabsPerPosition {
		numFuel := getFuelPart2(crabsPerPosition, targetPosition, costs)
		if numFuel < minFuel || minFuel < 0 {
			minFuel = numFuel
		}
	}
	return minFuel
}

func Run(path string) {
	initialPositions := utils.ParseStringAsIntList(utils.ReadFileAsString(path), ",")
	crabsPerPosition := getCrabsPerPosition(initialPositions)
	costs := getCostsPart2(crabsPerPosition)
	answer1 := doPart1(crabsPerPosition)
	answer2 := doPart2(crabsPerPosition, costs)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
