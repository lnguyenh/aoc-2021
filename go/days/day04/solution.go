package day04

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func doPart1(values []int) int {
	fmt.Println(values)
	return 0
}

func doPart2(values []int) int {
	fmt.Println(values)
	return 0
}

func Run(path string) {
	values := utils.ReadFileAsIntArray(path)
	answer1 := doPart1(values)
	answer2 := doPart2(values)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
