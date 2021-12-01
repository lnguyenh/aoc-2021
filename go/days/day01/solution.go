package day01

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/filereader"
	"github.com/lnguyenh/aoc-2021/utils"
)

func countIncreases(values []int, windowSize int) int {
	numValues := len(values)
	numBumps := 0
	for i := 0; i < numValues-windowSize; i++ {
		measurement1 := values[i : i+windowSize]
		measurement2 := values[i+1 : i+1+windowSize]
		if len(measurement2) == len(measurement2) &&
			utils.SumArray(measurement2) > utils.SumArray(measurement1) {
			numBumps += 1
		}
	}
	return numBumps
}

func Run(path string) {
	values := filereader.ReadAsIntArray(path)
	fmt.Printf("Part 1 answer: %v\n", countIncreases(values, 1))
	fmt.Printf("Part 2 answer: %v\n", countIncreases(values, 3))
}
