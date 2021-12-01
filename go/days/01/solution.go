package main

import (
	"fmt"
	"go/days/01/go/filereader"
	"go/days/01/go/utils"
)

func part1(path string) {
	values := filereader.ReadAsIntArray(path)

	var lastValue = -1
	var numBumps = 0
	for _, value := range values {
		if lastValue > 0 && value > lastValue {
			numBumps += 1
		}
		lastValue = value
	}
	fmt.Println(numBumps)
}

func part2(path string) {
	values := filereader.ReadAsIntArray(path)

	numBumps := 0
	numValues := len(values)

	for i := 0; i < numValues; i++ {
		measurement1 := values[i : i+3]
		measurement2 := values[i+1 : i+4]
		if len(measurement2) == len(measurement2) &&
			utils.SumArray(measurement2) > utils.SumArray(measurement1) {
			numBumps += 1
		}
	}
	fmt.Println(numBumps)
}

func main() {
	inputPath := "input/input.txt"
	part1(inputPath)
	part2(inputPath)
}
