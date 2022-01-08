package day24

import (
	"fmt"
)

func doPart1() int {
	doProgram([]int{3, 9, 9, 9, 9, 6, 9, 8, 7, 9, 9, 4, 2, 9})
	return 39999698799429
}

func doPart2() int {
	doProgram([]int{1, 8, 1, 1, 6, 1, 2, 1, 1, 3, 4, 1, 1, 7})
	return 18116121134117
}

func Run(path string) {
	answer1 := doPart1()
	answer2 := doPart2()
	fmt.Printf("Part 1 answer (purely found on a whiteboard by analyzing the input): %v\n", answer1)
	fmt.Printf("Part 2 answer (purely found on a whiteboard by analyzing the input): %v\n", answer2)
}
