package day09

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"sort"
)

type lowPoint struct {
	x      int
	y      int
	height int
	risk   int
}

func isNeighbourLower(grid [][]int, height, neighbourI, neighbourJ, minI, maxI, minJ, maxJ int) bool {
	return neighbourI >= minI &&
		neighbourI <= maxI &&
		neighbourJ >= minJ &&
		neighbourJ <= maxJ &&
		grid[neighbourI][neighbourJ] <= height
}

func getLowPoints(grid [][]int) []lowPoint {
	var lowPoints []lowPoint
	minI, maxI := 0, len(grid)-1
	minJ, maxJ := 0, len(grid[0])-1
	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			height := grid[i][j]
			if isNeighbourLower(grid, height, i-1, j, minI, maxI, minJ, maxJ) ||
				isNeighbourLower(grid, height, i+1, j, minI, maxI, minJ, maxJ) ||
				isNeighbourLower(grid, height, i, j-1, minI, maxI, minJ, maxJ) ||
				isNeighbourLower(grid, height, i, j+1, minI, maxI, minJ, maxJ) {
				continue
			}
			lowPoints = append(lowPoints, lowPoint{height: height, x: j, y: i, risk: height + 1})
		}
	}
	return lowPoints
}

func doPart1(grid [][]int) int {
	riskLevel := 0
	for _, point := range getLowPoints(grid) {
		riskLevel += point.risk
	}
	return riskLevel
}

func getEmptyGrid(grid [][]int) [][]int {
	length, width := len(grid), len(grid[0])
	emptyGrid := make([][]int, length)
	for i := range emptyGrid {
		emptyGrid[i] = make([]int, width)
	}
	return emptyGrid
}

func fillLowPoints(workGrid [][]int, lowPoints []lowPoint, basinSizes []int) {
	for i, point := range lowPoints {
		workGrid[point.y][point.x] = i + 1
		basinSizes[i] = 1
	}
}

func fillNines(workGrid [][]int, grid [][]int) {
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 9 {
				workGrid[i][j] = -1
			}
		}
	}
}

func expand(workGrid [][]int, basinSizes []int) int {
	minI, maxI := 0, len(workGrid)-1
	minJ, maxJ := 0, len(workGrid[0])-1
	numExpansions := 0

	for i := range workGrid {
		for j := range workGrid[0] {
			if workGrid[i][j] == 0 {
				expansions := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
				for _, expansionAttempt := range expansions {
					y, x := expansionAttempt[0], expansionAttempt[1]
					if y >= minI && y <= maxI && x >= minJ && x <= maxJ && workGrid[y][x] > 0 {
						basinId := workGrid[y][x]
						workGrid[i][j] = basinId
						basinSizes[basinId-1] += 1
						numExpansions++
						break
					}
				}
			}
		}
	}
	return numExpansions
}

func doPart2(grid [][]int) int {
	lowPoints := getLowPoints(grid)
	basinSizes := make([]int, len(lowPoints))
	workGrid := getEmptyGrid(grid)

	// Create a work grid where we will store the low-point-id that each point belongs to
	fillLowPoints(workGrid, lowPoints, basinSizes)
	fillNines(workGrid, grid)

	// Expand low points until nothing can be expanded
	numExpansions := 1
	for {
		if numExpansions == 0 {
			break
		}
		numExpansions = expand(workGrid, basinSizes)
	}

	// Generate answer
	sort.Ints(basinSizes)
	maxIndex := len(basinSizes) - 1
	return basinSizes[maxIndex] * basinSizes[maxIndex-1] * basinSizes[maxIndex-2]
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
