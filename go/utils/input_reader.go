package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileAsIntArray(path string) []int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var values []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}
	_ = file.Close()
	return values
}
