package day06

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func initializeCounterMap(day []int) map[int]int {
	counters := make(map[int]int)
	for _, counter := range day {
		counters[counter] += 1
	}
	return counters
}

func runOneMapDayOptimized(counters map[int]int) {
	numZeroes := counters[0]
	counters[0] = counters[1]
	counters[1] = counters[2]
	counters[2] = counters[3]
	counters[3] = counters[4]
	counters[4] = counters[5]
	counters[5] = counters[6]
	counters[6] = counters[7] + numZeroes
	counters[7] = counters[8]
	counters[8] = numZeroes
}

func doPart1(initialDay []int) int {
	counters := initializeCounterMap(initialDay)
	for i := 0; i < 80; i++ {
		runOneMapDayOptimized(counters)
	}
	return utils.SumIntMap(counters)
}

func doPart2(initialDay []int) int {
	counters := initializeCounterMap(initialDay)
	for i := 0; i < 256; i++ {
		runOneMapDayOptimized(counters)
	}
	return utils.SumIntMap(counters)
}

func Run(path string) {
	input := utils.ParseStringAsIntList(utils.ReadFileAsString(path), ",")
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
