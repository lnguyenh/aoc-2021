package utils

import (
	"strconv"
	"strings"
)

func ParseStringAsIntList(blob string, delimiter string) []int {
	var output []int
	for _, value := range strings.Split(blob, delimiter) {
		if value != "" {
			valueAsInt, _ := strconv.Atoi(value)
			output = append(output, valueAsInt)
		}
	}
	return output
}

func ParseStringAsStringList(blob string, delimiter string) []string {
	var output []string
	for _, value := range strings.Split(blob, delimiter) {
		if value != "" {
			output = append(output, value)
		}
	}
	return output
}

func ConvertBinaryStringToInt(text string) int {
	number, _ := strconv.ParseInt(text, 2, 64)
	return int(number)
}
