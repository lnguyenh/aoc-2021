package day15

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type locationPoint struct {
	risk       int
	x          int
	y          int
	neighbours []string
	minToEnd   *int
}

func getKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func createLocation(risk, x, y int) *locationPoint {
	neighbours := make([]string, 0, 4)
	minToEnd := -1
	return &locationPoint{
		risk:       risk,
		x:          x,
		y:          y,
		neighbours: neighbours,
		minToEnd:   &minToEnd,
	}
}

func buildLocations(locations [][]int, multiplicator int) map[string]*locationPoint {
	points := make(map[string]*locationPoint)
	miniLength := len(locations)
	miniWidth := len(locations[0])
	length := miniLength * multiplicator
	width := miniWidth * multiplicator

	for i := 0; i < multiplicator; i++ {
		for j := 0; j < multiplicator; j++ {
			for Y, row := range locations {
				for X := range row {
					x := X + (miniWidth * j)
					y := Y + (miniLength * i)

					risk := locations[Y][X]
					for z := 0; z < (i + j); z++ {
						risk = risk + 1
						if risk == 10 {
							risk = 1
						}
					}

					key := getKey(x, y)
					points[key] = createLocation(risk, x, y)

					potentialNeighbours := [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
					for _, potentialNeighbour := range potentialNeighbours {
						x0, y0 := potentialNeighbour[0], potentialNeighbour[1]
						if x0 >= 0 && x0 <= width-1 && y0 >= 0 && y0 <= length-1 {
							points[key].neighbours = append(points[key].neighbours, getKey(x0, y0))
						}
					}
				}
			}
		}
	}
	minToEnd := points[getKey(length-1, width-1)].risk
	points[getKey(length-1, width-1)].minToEnd = &minToEnd
	return points
}

func populateMinRisk(points *map[string]*locationPoint) {
	numPopulated := -1
	for {
		if numPopulated == 0 {
			break
		}
		numPopulated = 0
		for _, pointP := range *points {
			for _, neighbourKey := range pointP.neighbours {
				neighbourP := (*points)[neighbourKey]
				if *neighbourP.minToEnd > 0 {
					minToEndCandidate := pointP.risk + *neighbourP.minToEnd
					if *pointP.minToEnd > minToEndCandidate || *pointP.minToEnd < 0 {
						pointP.minToEnd = &minToEndCandidate
						numPopulated++
					}
				}
			}
		}
	}
}

func printGrid(points map[string]*locationPoint) {
	widthAndLength := 0
	for _, point := range points {
		if point.x > widthAndLength {
			widthAndLength = point.x
		}
	}
	for i := 0; i < widthAndLength; i++ {
		for j := 0; j < widthAndLength; j++ {
			fmt.Printf("%v ", points[getKey(j, i)].risk)
		}
		fmt.Printf("\n")
	}
}

func getMinRisk(points map[string]*locationPoint) int {
	populateMinRisk(&points)
	startPoint := points[getKey(0, 0)]
	return *startPoint.minToEnd - startPoint.risk
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	points := buildLocations(input, 1)
	points2 := buildLocations(input, 5)
	answer1 := getMinRisk(points)
	answer2 := getMinRisk(points2)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
