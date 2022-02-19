package day15

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"math"
)

var maxValue = math.MaxUint32

type locationPoint struct {
	risk       int
	x          int
	y          int
	neighbours []string
	minToEnd   int
}

func getKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func createPoint(risk, x, y int) *locationPoint {
	neighbours := make([]string, 0, 4)
	return &locationPoint{
		risk:       risk,
		x:          x,
		y:          y,
		neighbours: neighbours,
		minToEnd:   maxValue,
	}
}

func buildPoints(rawInput [][]int, multiplicator int) (map[string]*locationPoint, int, int) {
	points := make(map[string]*locationPoint)
	baseLength := len(rawInput)
	baseWidth := len(rawInput[0])
	length := baseLength * multiplicator
	width := baseWidth * multiplicator

	for i := 0; i < multiplicator; i++ {
		for j := 0; j < multiplicator; j++ {
			for Y, row := range rawInput {
				for X := range row {
					x := X + (baseWidth * j)
					y := Y + (baseLength * i)

					risk := rawInput[Y][X]
					for z := 0; z < (i + j); z++ {
						risk = risk + 1
						if risk == 10 {
							risk = 1
						}
					}

					key := getKey(x, y)
					points[key] = createPoint(risk, x, y)

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

	// Set the only known minToEnd: the last point
	endPointKey := getKey(length-1, width-1)
	points[endPointKey].minToEnd = points[endPointKey].risk
	return points, length, width
}

func populateMinRisk(points *map[string]*locationPoint) {
	numPopulated := -1
	for {
		if numPopulated == 0 {
			break
		}
		numPopulated = 0
		for _, point := range *points {
			for _, neighbourKey := range point.neighbours {
				neighbour := (*points)[neighbourKey]
				if neighbour.minToEnd != maxValue {
					minToEndCandidate := point.risk + neighbour.minToEnd
					if point.minToEnd > minToEndCandidate {
						point.minToEnd = minToEndCandidate
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
	return startPoint.minToEnd - startPoint.risk
}

func runDjikstra(points map[string]*locationPoint, startNode string) (map[string]string, map[string]int) {
	unvisitedNodes := make(map[string]bool)
	for node := range points {
		unvisitedNodes[node] = true
	}

	shortestPath := make(map[string]int)
	for node := range points {
		shortestPath[node] = maxValue
	}
	shortestPath[startNode] = 0

	previousNodes := make(map[string]string)

	for {
		if len(unvisitedNodes) == 0 {
			break
		}

		// Find node with lowest value
		currentMinNode := ""
		for node := range unvisitedNodes {
			if currentMinNode == "" {
				currentMinNode = node
			} else if shortestPath[node] < shortestPath[currentMinNode] {
				currentMinNode = node
			}
		}

		for _, neighbour := range points[currentMinNode].neighbours {
			tentativeValue := shortestPath[currentMinNode] + points[neighbour].risk
			if tentativeValue < shortestPath[neighbour] {
				shortestPath[neighbour] = tentativeValue
				previousNodes[neighbour] = currentMinNode
			}
		}

		delete(unvisitedNodes, currentMinNode)
	}

	return previousNodes, shortestPath
}

func getMinRiskDjikstra(points map[string]*locationPoint, length int, width int) int {
	_, shortestPath := runDjikstra(points, getKey(0, 0))
	return shortestPath[getKey(length-1, width-1)]
}

func getSolution(rawInput [][]int, multiplicator int, doPrintGrid bool, useDjikstra bool) int {
	points, length, width := buildPoints(rawInput, multiplicator)
	if doPrintGrid {
		printGrid(points)
	}
	if useDjikstra {
		return getMinRiskDjikstra(points, length, width)
	} else {
		return getMinRisk(points)
	}
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	useDjikstra := false
	doPrintGrid := false
	answer1 := getSolution(input, 1, doPrintGrid, useDjikstra)
	answer2 := getSolution(input, 5, doPrintGrid, useDjikstra)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
