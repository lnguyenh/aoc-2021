package day15

import (
	"fmt"
	"math"
)

// TODO remove this file once demo-ed

var maxValue = math.MaxUint32

// Original solution using bruteforce "contagion"
func populateMinRisk(nodes *map[string]*locationNode, length, width int) {
	// Initialize minToEnd
	for _, node := range *nodes {
		node.minToEnd = maxValue
	}
	// Set the only known minToEnd: the last node
	endNodeKey := getNodeName(length-1, width-1)
	(*nodes)[endNodeKey].minToEnd = (*nodes)[endNodeKey].risk

	numPopulated := -1
	for {
		if numPopulated == 0 {
			break
		}
		numPopulated = 0
		for _, node := range *nodes {
			for _, neighbourKey := range node.neighbours {
				neighbour := (*nodes)[neighbourKey]
				if neighbour.minToEnd != maxValue {
					minToEndCandidate := node.risk + neighbour.minToEnd
					if node.minToEnd > minToEndCandidate {
						node.minToEnd = minToEndCandidate
						numPopulated++
					}
				}
			}
		}
	}
}

func getMinRisk(nodes map[string]*locationNode, length, width int) int {
	populateMinRisk(&nodes, length, width)
	startNode := nodes[getNodeName(0, 0)]
	return startNode.minToEnd - startNode.risk
}

// First version of Djikstra, unoptimized
func runDjikstra(nodes map[string]*locationNode, startNode, stopNode string) int {
	unvisitedNodes := make(map[string]bool)
	for node := range nodes {
		unvisitedNodes[node] = true
	}

	shortestPath := make(map[string]int)
	for node := range nodes {
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

		for _, neighbour := range nodes[currentMinNode].neighbours {
			tentativeValue := shortestPath[currentMinNode] + nodes[neighbour].risk
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

func printGrid(nodes map[string]*locationNode) {
	widthAndLength := 0
	for _, node := range nodes {
		if node.x > widthAndLength {
			widthAndLength = node.x
		}
	}
	for i := 0; i < widthAndLength; i++ {
		for j := 0; j < widthAndLength; j++ {
			fmt.Printf("%v ", nodes[getNodeName(j, i)].risk)
		}
		fmt.Printf("\n")
	}
}
