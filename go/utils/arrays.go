package utils

import (
	"strconv"
)

func SumArray(measurement []int) int {
	sum := 0
	for _, value := range measurement {
		sum += value
	}
	return sum
}

func ConvertBinaryStringToInt(text string) int {
	number, _ := strconv.ParseInt(text, 2, 64)
	return int(number)
}
