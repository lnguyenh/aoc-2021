package day03

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func getMostCommonInColumn(report [][]rune, column int) rune {
	var numOnes, numZeroes int
	for _, value := range report {
		if value[column] == '0' {
			numZeroes += 1
		} else {
			numOnes += 1
		}
	}
	mostCommon := '0'
	if numOnes >= numZeroes {
		mostCommon = '1'
	}
	return mostCommon
}

func getLeastCommonInColumn(report [][]rune, column int) rune {
	leastCommon := '0'
	if getMostCommonInColumn(report, column) == '0' {
		leastCommon = '1'
	}
	return leastCommon
}

func reduceReport(report [][]rune, bitCriteria rune, column int) [][]rune {
	var reducedReport [][]rune
	for _, value := range report {
		if value[column] == bitCriteria {
			reducedReport = append(reducedReport, value)
		}
	}
	return reducedReport
}

func getOxygen(report [][]rune) string {
	reducedReport := report
	column := 0
	for ok := true; ok; ok = len(reducedReport) > 1 {
		reducedReport = reduceReport(reducedReport, getMostCommonInColumn(reducedReport, column), column)
		column += 1
	}
	return string(reducedReport[0])
}

func getCo2(report [][]rune) string {
	reducedReport := report
	column := 0
	for ok := true; ok; ok = len(reducedReport) > 1 {
		reducedReport = reduceReport(reducedReport, getLeastCommonInColumn(reducedReport, column), column)
		column += 1
	}
	return string(reducedReport[0])
}

func doPart1(report [][]rune) int {
	var gamma, epsilon []rune
	for column := range report[0] {
		if getMostCommonInColumn(report, column) == '1' {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}
	gammaAsInt := utils.ConvertBinaryStringToInt(string(gamma))
	epsilonAsInt := utils.ConvertBinaryStringToInt(string(epsilon))
	return gammaAsInt * epsilonAsInt
}

func doPart2(report [][]rune) int {
	oxygen := getOxygen(report)
	co2 := getCo2(report)
	oxygenAsInt := utils.ConvertBinaryStringToInt(oxygen)
	co2AsInt := utils.ConvertBinaryStringToInt(co2)
	return oxygenAsInt * co2AsInt
}

func Run(path string) {
	report := utils.ReadFileAsSliceOfRuneSlices(path)
	answer1 := doPart1(report)
	answer2 := doPart2(report)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
