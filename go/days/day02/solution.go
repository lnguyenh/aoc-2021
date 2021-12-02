package day02

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strconv"
)

func applyOneInstruction(start []int, instruction []string) []int {
	moveValue, _ := strconv.Atoi(instruction[1])
	var stop []int
	switch instruction[0] {
	case "forward":
		stop = []int{start[0] + moveValue, start[1]}
	case "down":
		stop = []int{start[0], start[1] - moveValue}
	case "up":
		stop = []int{start[0], start[1] + moveValue}
	}
	return stop
}

func applyOneAimInstruction(start []int, aim int, instruction []string) ([]int, int) {
	changeValue, _ := strconv.Atoi(instruction[1])
	var stop []int = start
	var newAim int = aim
	switch instruction[0] {
	case "forward":
		stop = []int{start[0] + changeValue, start[1] + changeValue*aim}
	case "down":
		newAim += changeValue
	case "up":
		newAim -= changeValue
	}
	return stop, newAim
}

func doPart1(instructions [][]string) int {
	position := []int{0, 0}
	for _, instruction := range instructions {
		position = applyOneInstruction(position, instruction)
	}
	result := position[0] * -position[1]
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
