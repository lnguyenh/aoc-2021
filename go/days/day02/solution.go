package day02

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strconv"
)

func applyOneInstruction(position []int, instruction []string) []int {
	moveValue, _ := strconv.Atoi(instruction[1])
	var newPosition []int
	switch instruction[0] {
	case "forward":
		newPosition = []int{position[0] + moveValue, position[1]}
	case "down":
		newPosition = []int{position[0], position[1] + moveValue}
	case "up":
		newPosition = []int{position[0], position[1] - moveValue}
	}
	return newPosition
}

func applyOneAimInstruction(position []int, aim int, instruction []string) ([]int, int) {
	changeValue, _ := strconv.Atoi(instruction[1])
	var newPosition = position
	var newAim = aim
	switch instruction[0] {
	case "forward":
		newPosition = []int{position[0] + changeValue, position[1] + changeValue*aim}
	case "down":
		newAim += changeValue
	case "up":
		newAim -= changeValue
	}
	return newPosition, newAim
}

func doPart1(instructions [][]string) int {
	position := []int{0, 0}
	for _, instruction := range instructions {
		position = applyOneInstruction(position, instruction)
	}
	result := position[0] * position[1]
	return result
}

func doPart2(instructions [][]string) int {
	position := []int{0, 0}
	aim := 0
	for _, instruction := range instructions {
		position, aim = applyOneAimInstruction(position, aim, instruction)
	}
	result := position[0] * position[1]
	return result
}

func Run(path string) {
	values := utils.ReadFileAsArrayOfStringArrays(path)
	answer1 := doPart1(values)
	answer2 := doPart2(values)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
