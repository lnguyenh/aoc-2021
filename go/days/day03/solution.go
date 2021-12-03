package day03

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func hasColumnMoreOnes(report [][]string, column int) bool {
	var numOnes, numZeroes int
	for _, value := range report {
		if value[column] == "0" {
			numZeroes += 1
		} else {
			numOnes += 1
		}
	}
	return numOnes > numZeroes
}

func doPart1(report [][]string) int {
	var gamma, epsilon []rune
	var gammaNumber, epsilonNumber int
	for column := range report[0] {
		if hasColumnMoreOnes(report, column) {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}
	gammaNumber = utils.ConvertBinaryStringToInt(string(gamma))
	epsilonNumber = utils.ConvertBinaryStringToInt(string(epsilon))
	return gammaNumber * epsilonNumber
}

func doPart2(report [][]string) int {
	fmt.Println(report)
	return 0
}

func Run(path string) {
	report := utils.ReadFileAsArrayOfStringArrays(path, "")
	answer1 := doPart1(report)
	answer2 := doPart2(report)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
