package day03

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

func getMostCommonInColumn(report [][]string, column int) string {
	var numOnes, numZeroes int
	for _, value := range report {
		if value[column] == "0" {
			numZeroes += 1
		} else {
			numOnes += 1
		}
	}

	mostCommon := "0"
	if numOnes >= numZeroes {
		mostCommon = "1"
	}
	return mostCommon
}

func getLeastCommonInColumn(report [][]string, column int) string {
	leastCommon := "0"
	if getMostCommonInColumn(report, column) == "0" {
		leastCommon = "1"
	}
	return leastCommon
}

func reduceReport(report [][]string, bitCriteria string, column int) [][]string {
	var reducedReport [][]string
	for _, value := range report {
		if value[column] == bitCriteria {
			reducedReport = append(reducedReport, value)
		}
	}
	return reducedReport
}

func getOxygen(report [][]string) string {
	reducedReport := report
	column := 0
	for ok := true; ok; ok = len(reducedReport) > 1 {
		reducedReport = reduceReport(reducedReport, getMostCommonInColumn(reducedReport, column), column)
		column += 1
	}
	return strings.Join(reducedReport[0], "")
}

func getCo2(report [][]string) string {
	reducedReport := report
	column := 0
	for ok := true; ok; ok = len(reducedReport) > 1 {
		reducedReport = reduceReport(reducedReport, getLeastCommonInColumn(reducedReport, column), column)
		column += 1
	}
	return strings.Join(reducedReport[0], "")
}

func doPart1(report [][]string) int {
	var gamma, epsilon []rune
	for column := range report[0] {
		if getMostCommonInColumn(report, column) == "1" {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}
	gammaNumber := utils.ConvertBinaryStringToInt(string(gamma))
	epsilonNumber := utils.ConvertBinaryStringToInt(string(epsilon))
	return gammaNumber * epsilonNumber
}

func doPart2(report [][]string) int {
	oxygen := getOxygen(report)
	co2 := getCo2(report)
	oxygenNumber := utils.ConvertBinaryStringToInt(oxygen)
	co2Number := utils.ConvertBinaryStringToInt(co2)
	return oxygenNumber * co2Number
}

func Run(path string) {
	report := utils.ReadFileAsArrayOfStringArrays(path, "")
	answer1 := doPart1(report)
	answer2 := doPart2(report)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
