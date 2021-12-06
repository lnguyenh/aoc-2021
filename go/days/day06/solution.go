package day06

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

var normalDelay = 6
var firstDelay = 8

func runOneDay(day []int) []int {
	newDay := make([]int, 0, len(day))
	for _, counter := range day {
		if counter == 0 {
			newDay = append(newDay, normalDelay)
			newDay = append(newDay, firstDelay)
		} else {
			newDay = append(newDay, counter-1)
		}
	}
	return newDay
}

func doPart1(initialDay []int) int {
	day := initialDay
	for i := 0; i < 80; i++ {
		day = runOneDay(day)
	}
	return len(day)
}

func initializeCounterMap(day []int) map[int]int {
	counters := make(map[int]int)
	for _, counter := range day {
		counters[counter] += 1
	}
	return counters
}

func runOneMapDay(counters map[int]int) map[int]int {
	newCounters := make(map[int]int)
	for daysLeft, numFish := range counters {
		if daysLeft == 0 {
			newCounters[normalDelay] += numFish
			newCounters[firstDelay] = numFish
		} else {
			newCounters[daysLeft-1] += numFish
		}
	}
	return newCounters
}

func doPart2(initialDay []int) int {
	counters := initializeCounterMap(initialDay)
	for i := 0; i < 256; i++ {
		counters = runOneMapDay(counters)
	}
	return utils.SumIntMap(counters)
}

func Run(path string) {
	input := utils.ParseStringAsIntList(utils.ReadFileAsStringSlice(path, "\n")[0], ",")
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
