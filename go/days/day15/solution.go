package day15

import (
	"container/heap"
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type locationNode struct {
	risk       int
	x          int
	y          int
	neighbours []string
	minToEnd   int // TODO: remove this legacy field when not needed for demo anymore
}

func getNodeName(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func createNode(risk, x, y int) *locationNode {
	neighbours := make([]string, 0, 4)
	return &locationNode{
		risk:       risk,
		x:          x,
		y:          y,
		neighbours: neighbours,
	}
}

func createNodes(rawInput [][]int, multiplicator int) (map[string]*locationNode, int, int) {
	nodes := make(map[string]*locationNode)
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
					nodes[key] = createNode(risk, x, y)

					potentialNeighbours := [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
					for _, potentialNeighbour := range potentialNeighbours {
						x0, y0 := potentialNeighbour[0], potentialNeighbour[1]
						if x0 >= 0 && x0 <= width-1 && y0 >= 0 && y0 <= length-1 {
							nodes[key].neighbours = append(nodes[key].neighbours, getNodeName(x0, y0))
						}
					}
				}
			}
		}
	}
	return nodes, length, width
}

func runDjikstraHeap(nodes map[string]*locationNode, startNode, stopNode string) int {
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

		for _, neighbourName := range nodes[currentMinNode.name].neighbours {
			if !visitedNodes[neighbourName] {
				neighbourTotalRisk := currentMinNode.totalRisk + nodes[neighbourName].risk
				heap.Push(&minHeap, heapNode{name: neighbourName, totalRisk: neighbourTotalRisk})
			}
		}

		visitedNodes[currentMinNode.name] = true
	}
}

func getSolution(rawInput [][]int, multiplicator int, doPrintGrid bool) int {
	nodes, length, width := createNodes(rawInput, multiplicator)
	if doPrintGrid {
		printGrid(nodes)
	}
	return runDjikstraHeap(nodes, getNodeName(0, 0), getNodeName(length-1, width-1))
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	answer1 := getSolution(input, 1, false)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := getSolution(input, 5, false)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
