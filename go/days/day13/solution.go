package day13

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strconv"
	"strings"
)

type instructionFold struct {
	axe   string
	value int
}

func getKey(point [2]int) string {
	return fmt.Sprintf("%v-%v", point[0], point[1])
}

func applyFold(points [][2]int, fold instructionFold) [][2]int {
	pointsAfterFold := make([][2]int, 0, len(points))
	exists := map[string]bool{}
	for _, point := range points {
		x, y := point[0], point[1]
		var newPoint [2]int

		switch fold.axe {
		case "x":
			xAxe := fold.value
			if x < xAxe {
				newPoint[0], newPoint[1] = x, y
			} else if x > xAxe {
				newPoint[0], newPoint[1] = xAxe-(x-xAxe), y
			}
		case "y":
			yAxe := fold.value
			if y < yAxe {
				newPoint[0], newPoint[1] = x, y
			} else if y > yAxe {
				newPoint[0], newPoint[1] = x, yAxe-(y-yAxe)
			}
		}
		if !exists[getKey(newPoint)] {
			pointsAfterFold = append(pointsAfterFold, newPoint)
			exists[getKey(newPoint)] = true
		}
	}
	return pointsAfterFold

}

func printGrid(points [][2]int) {
	var maxX, maxY int
	exists := map[string]bool{}

	// Get boundaries and existence map
	for _, point := range points {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
		exists[getKey(point)] = true
	}

	// Print
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if exists[getKey([2]int{j, i})] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

func doPart1(points [][2]int, folds []instructionFold) int {
	newPoints := applyFold(points, folds[0])
	return len(newPoints)
}

func doPart2(points [][2]int, folds []instructionFold) int {
	newPoints := applyFold(points, folds[0])
	for _, fold := range folds[1:] {
		newPoints = applyFold(newPoints, fold)
	}
	printGrid(newPoints)
	return len(newPoints)
}

func parseInput(input []string) ([][2]int, []instructionFold) {
	points := make([][2]int, 0, 1000)
	for _, line := range strings.Split(input[0], "\n") {
		if line != "" {
			xy := strings.Split(line, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			points = append(points, [2]int{x, y})
		}
	}
	var folds []instructionFold
	for _, rawInstruction := range input[1:] {
		parts := strings.Split(rawInstruction, "=")
		value, _ := strconv.Atoi(strings.TrimSuffix(parts[1], "\n"))
		folds = append(folds, instructionFold{axe: parts[0], value: value})
	}
	return points, folds

}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "fold along ")
	points, folds := parseInput(input)
	answer1 := doPart1(points, folds)
	answer2 := doPart2(points, folds)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
