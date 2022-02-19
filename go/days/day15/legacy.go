package day15

import (
	"fmt"
	"math"
)

var maxValue = math.MaxUint32

func populateMinRisk(points *map[string]*locationPoint, length, width int) {
	// Initialize minToEnd
	for _, point := range *points {
		point.minToEnd = maxValue
	}
	// Set the only known minToEnd: the last point
	endPointKey := getNodeName(length-1, width-1)
	(*points)[endPointKey].minToEnd = (*points)[endPointKey].risk

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

func getMinRisk(points map[string]*locationPoint, length, width int) int {
	populateMinRisk(&points, length, width)
	startPoint := points[getNodeName(0, 0)]
	return startPoint.minToEnd - startPoint.risk
}

func runDjikstra(points map[string]*locationPoint, startNode, stopNode string) int {
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

		if currentMinNode == stopNode {
			break
		}

		delete(unvisitedNodes, currentMinNode)
	}

	return shortestPath[stopNode]
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
			fmt.Printf("%v ", points[getNodeName(j, i)].risk)
		}
		fmt.Printf("\n")
	}
}
