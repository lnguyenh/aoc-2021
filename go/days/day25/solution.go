package day25

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type aocSea struct {
	cucumbers map[aocCoordinates]rune
	width     int
	length    int
}

type aocCoordinates struct {
	x int
	y int
}

func (s *aocSea) print() {
	for i := 0; i < s.length; i++ {
		for j := 0; j < s.width; j++ {
			cucumber, exists := s.cucumbers[aocCoordinates{j, i}]
			if exists {
				fmt.Printf("%c", cucumber)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (s *aocSea) doStep() int {
	numMoves := 0
	newCucumbers := make(map[aocCoordinates]rune)
	newCucumbers2 := make(map[aocCoordinates]rune)
	for coordinates, cucumber := range s.cucumbers {
		switch cucumber {
		case '>':
			destination := aocCoordinates{(coordinates.x + 1) % s.width, coordinates.y}
			_, exists := s.cucumbers[destination]
			if !exists {
				newCucumbers[destination] = '>'
				numMoves++
			} else {
				newCucumbers[coordinates] = '>'
			}
		case 'v':
			newCucumbers[coordinates] = 'v'
		}
	}
	for coordinates, cucumber := range newCucumbers {
		switch cucumber {
		case '>':
			newCucumbers2[coordinates] = '>'
		case 'v':
			destination := aocCoordinates{coordinates.x, (coordinates.y + 1) % s.length}
			_, exists := newCucumbers[destination]
			if !exists {
				newCucumbers2[destination] = 'v'
				numMoves++
			} else {
				newCucumbers2[coordinates] = 'v'
			}
		}
	}
	s.cucumbers = newCucumbers2
	return numMoves
}

func doPart1(input [][]rune) int {
	sea := aocSea{cucumbers: make(map[aocCoordinates]rune), width: len(input[0]), length: len(input)}
	for i := range input {
		for j := range input[i] {
			if input[i][j] != '.' {
				sea.cucumbers[aocCoordinates{x: j, y: i}] = input[i][j]
			}
		}
	}

	numMoves := 1
	numSteps := 0
	for {
		if numMoves == 0 {
			break
		}
		numMoves = sea.doStep()
		numSteps++
	}

	return numSteps
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfRuneSlices(path)
	answer1 := doPart1(input)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
