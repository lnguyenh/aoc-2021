package day15

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type locationPoint struct {
	risk       int
	isEnd      bool
	x          int
	y          int
	key        string
	neighbours []string
}

func getKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func createLocation(risk, x, y int) *locationPoint {
	neighbours := make([]string, 0, 4)
	return &locationPoint{
		risk:       risk,
		isEnd:      false,
		x:          x,
		y:          y,
		key:        getKey(x, y),
		neighbours: neighbours,
	}
}

func buildLocations(locations [][]int) map[string]*locationPoint {
	points := make(map[string]*locationPoint)
	length := len(locations)
	width := len(locations[0])

	for y, row := range locations {
		for x := range row {
			key := getKey(x, y)
			points[key] = createLocation(locations[y][x], x, y)

			potentialNeighbours := [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
			for _, potentialNeighbour := range potentialNeighbours {
				x0, y0 := potentialNeighbour[0], potentialNeighbour[1]
				if x0 >= 0 && x0 <= width-1 && y0 >= 0 && y0 <= length-1 {
					points[key].neighbours = append(points[key].neighbours, getKey(x0, y0))
				}
			}
		}
	}
	points[getKey(length-1, width-1)].isEnd = true
	return points
}

func getMinRisk(points *map[string]*locationPoint, pointKey string, riskAccumulator int, currentTrail []string, currentMin *int, currentBestTrail *[]string) {
	point := (*points)[pointKey]
	accumulator := riskAccumulator + point.risk
	trail := append(currentTrail, point.key)

	if point.isEnd {
		if *currentMin < 0 || accumulator < *currentMin {
			*currentMin = accumulator
			currentBestTrail = &trail
		}
		return
	}

	if *currentMin > 0 && accumulator >= *currentMin {
		return
	}

	candidates := utils.StringsNotInSlice(point.neighbours, trail)
	if len(candidates) == 0 {
		return
	}
	for _, neighbourKey := range candidates {
		getMinRisk(points, neighbourKey, accumulator, trail, currentMin, currentBestTrail)
	}

}

func doPart1(points map[string]*locationPoint) int {
	startPoint := points[getKey(0, 0)]
	minRisk := -1
	var bestTrail *[]string
	getMinRisk(&points, startPoint.key, 0, make([]string, 0), &minRisk, bestTrail)
	return minRisk - startPoint.risk
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	points := buildLocations(input)
	answer1 := doPart1(points)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
