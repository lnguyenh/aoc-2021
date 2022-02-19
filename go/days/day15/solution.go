package day15

import (
	"container/heap"
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type locationPoint struct {
	risk       int
	x          int
	y          int
	neighbours []string
	minToEnd   int // TODO Legacy remove
}

func getNodeName(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func createPoint(risk, x, y int) *locationPoint {
	neighbours := make([]string, 0, 4)
	return &locationPoint{
		risk:       risk,
		x:          x,
		y:          y,
		neighbours: neighbours,
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

					key := getNodeName(x, y)
					points[key] = createPoint(risk, x, y)

					potentialNeighbours := [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
					for _, potentialNeighbour := range potentialNeighbours {
						x0, y0 := potentialNeighbour[0], potentialNeighbour[1]
						if x0 >= 0 && x0 <= width-1 && y0 >= 0 && y0 <= length-1 {
							points[key].neighbours = append(points[key].neighbours, getNodeName(x0, y0))
						}
					}
				}
			}
		}
	}
	return points, length, width
}

func runDjikstraHeap(points map[string]*locationPoint, startNode, stopNode string) int {
	visitedNodes := make(map[string]bool)

	// used to store the "active" minimum distances from start node
	minHeap := aocHeap{heapNode{startNode, 0}}

	for {
		currentMinNode := heap.Pop(&minHeap).(heapNode)

		if visitedNodes[currentMinNode.name] {
			continue
		}

		if currentMinNode.name == stopNode {
			return currentMinNode.totalRisk
		}

		for _, neighbourName := range points[currentMinNode.name].neighbours {
			if !visitedNodes[neighbourName] {
				neighbourTotalRisk := currentMinNode.totalRisk + points[neighbourName].risk
				heap.Push(&minHeap, heapNode{name: neighbourName, totalRisk: neighbourTotalRisk})
			}
		}

		visitedNodes[currentMinNode.name] = true
	}
}

func getSolution(rawInput [][]int, multiplicator int, doPrintGrid bool) int {
	points, length, width := buildPoints(rawInput, multiplicator)
	if doPrintGrid {
		printGrid(points)
	}
	return runDjikstraHeap(points, getNodeName(0, 0), getNodeName(length-1, width-1))
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	answer1 := getSolution(input, 1, false)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := getSolution(input, 5, false)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
