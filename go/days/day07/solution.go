package day07

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func getNumCrabsPerPosition(positions []int) []int {
	maxX := utils.MaxSlice(positions)
	numCrabsPerPositionArray := make([]int, maxX+1)
	for _, position := range positions {
		numCrabsPerPositionArray[position] += 1
	}
	return numCrabsPerPositionArray
}

func getFuelPart1(numCrabsPerPosition []int, targetPosition int) int {
	totalFuel := 0
	for position, numCrabs := range numCrabsPerPosition {
		numMoves := numCrabs * utils.IntAbs(targetPosition-position)
		totalFuel += numMoves
	}
	return totalFuel
}

func getCostsPart2(numCrabsPerPosition []int) []int {
	costs := make([]int, len(numCrabsPerPosition)*utils.MaxSlice(numCrabsPerPosition))
	accumulator := 0
	for numSteps := range costs {
		costs[numSteps] = numSteps + accumulator
		accumulator += numSteps
	}
	return costs
}

func getFuelPart2(numCrabsPerPosition []int, targetPosition int, costs []int) int {
	totalFuel := 0
	for position, numCrabs := range numCrabsPerPosition {
		numMoves := utils.IntAbs(targetPosition - position)
		totalFuel += costs[numMoves] * numCrabs
	}
	return totalFuel
}

func doPart1(numCrabsPerPosition []int) int {
	minMoves := -1
	for targetPosition := range numCrabsPerPosition {
		numMoves := getFuelPart1(numCrabsPerPosition, targetPosition)
		if numMoves < minMoves || minMoves < 0 {
			minMoves = numMoves
		}
	}
	return minMoves
}

func doPart2(numCrabsPerPosition []int, costs []int) int {
	minFuel := -1
	for targetPosition := range numCrabsPerPosition {
		numFuel := getFuelPart2(numCrabsPerPosition, targetPosition, costs)
		if numFuel < minFuel || minFuel < 0 {
			minFuel = numFuel
		}
	}
	return minFuel
}

func Run(path string) {
	initialPositions := utils.ParseStringAsIntList(utils.ReadFileAsString(path), ",")
	numCrabsPerPosition := getNumCrabsPerPosition(initialPositions)
	costs := getCostsPart2(numCrabsPerPosition)
	answer1 := doPart1(numCrabsPerPosition)
	answer2 := doPart2(numCrabsPerPosition, costs)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
