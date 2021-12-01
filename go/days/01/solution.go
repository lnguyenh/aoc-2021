package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	file, _ := os.Open("input/input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	var lastValue = -1
	var numBumps = 0
	for _, line := range lines {
		currentValue, _ := strconv.Atoi(line)
		if lastValue > 0 && currentValue > lastValue {
			numBumps += 1
		}
		lastValue, _ = strconv.Atoi(line)
	}
	fmt.Println(numBumps)
}

func sumArray(measurement []int) int {
	sum := 0
	for _, value := range measurement {
		sum += value
	}
	return sum
}

func part2() {
	file, _ := os.Open("input/input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var values []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}
	file.Close()

	numBumps := 0
	numValues := len(values)

	for i := 0; i < numValues; i++ {
		measurement1 := values[i : i+3]
		measurement2 := values[i+1 : i+4]
		if len(measurement2) == len(measurement2) && sumArray(measurement2) > sumArray(measurement1) {
			numBumps += 1
		}
	}
	fmt.Println(numBumps)
}

func main() {
	part1()
	part2()
}
