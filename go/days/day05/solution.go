package day05

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
	"strconv"
)

type ventVector struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (vector ventVector) getPoints(useDiagonal bool) [][]int {
	var points [][]int
	var originalPoint = []int{vector.x1, vector.y1}
	var x, y int
	if vector.x1 != vector.x2 && vector.y1 != vector.y2 {
		if useDiagonal {
			points = append(points, originalPoint)
			x = vector.x1
			y = vector.y1
			for {
				if x == vector.x2 && y == vector.y2 {
					break
				}
				if x < vector.x2 {
					x = x + 1
				} else if x > vector.x2 {
					x = x - 1
				}
				if y < vector.y2 {
					y = y + 1
				} else if y > vector.y2 {
					y = y - 1
				}
				var newPoint = []int{x, y}
				points = append(points, newPoint)
			}
		}
	} else {
		points = append(points, originalPoint)
		x = vector.x1
		y = vector.y1
		for {
			if x == vector.x2 && y == vector.y2 {
				break
			}
			if x < vector.x2 {
				x = x + 1
			} else if x > vector.x2 {
				x = x - 1
			} else if y < vector.y2 {
				y = y + 1
			} else if y > vector.y2 {
				y = y - 1
			}
			var newPoint = []int{x, y}
			points = append(points, newPoint)
		}
	}
	return points
}

func createVector(blob string) ventVector {
	r := regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)
	matches := r.FindAllStringSubmatch(blob, -1)
	var coordinates []int
	for i, match := range matches[0] {
		if i < 1 {
			continue
		}
		number, _ := strconv.Atoi(match)
		coordinates = append(coordinates, number)
	}
	newVector := ventVector{
		x1: coordinates[0],
		y1: coordinates[1],
		x2: coordinates[2],
		y2: coordinates[3],
	}
	return newVector
}

func createVectors(blobs []string) []ventVector {
	var vectors []ventVector
	for _, blob := range blobs {
		vectors = append(vectors, createVector(blob))
	}
	return vectors
}

func createGrid(I, J int) [][]int {
	grid := make([][]int, I)
	for i := 0; i < I; i++ {
		grid[i] = make([]int, J)
	}
	return grid
}

func applyVector(vector ventVector, grid [][]int, useDiagonal bool) {
	points := vector.getPoints(useDiagonal)
	for _, point := range points {
		grid[point[1]][point[0]] += 1
	}
}

func countOverlaps(grid [][]int) int {
	count := 0
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[i][j] >= 2 {
				count++
			}
		}
	}
	return count
}

func printGrid(grid [][]int, maxX int, maxY int) {
	for i := 0; i < maxY; i++ {
		fmt.Println(grid[i][:maxX])
	}
}

func doPart1(vectors []ventVector) int {
	grid := createGrid(1000, 1000)
	for _, v := range vectors {
		applyVector(v, grid, false)
	}
	return countOverlaps(grid)
}

func doPart2(vectors []ventVector) int {
	grid := createGrid(1000, 1000)
	for _, v := range vectors {
		applyVector(v, grid, true)
	}
	return countOverlaps(grid)
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	vectors := createVectors(input)

	answer1 := doPart1(vectors)
	answer2 := doPart2(vectors)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
