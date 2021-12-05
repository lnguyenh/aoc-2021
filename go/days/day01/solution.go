package day01

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func countIncreases(values []int, windowSize int) int {
	numValues := len(values)
	numBumps := 0
	for i := 0; i < numValues-windowSize; i++ {
		measurement1 := values[i : i+windowSize]
		measurement2 := values[i+1 : i+1+windowSize]
		if utils.SumSlice(measurement2) > utils.SumSlice(measurement1) {
			numBumps += 1
		}
	}
	return numBumps
}

func Run(path string) {
	values := utils.ReadFileAsIntSlice(path)
	answer1 := countIncreases(values, 1)
	answer2 := countIncreases(values, 3)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
