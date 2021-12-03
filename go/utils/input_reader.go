package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsIntArray(path string) []int {
	var values []int
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}
	_ = file.Close()
	return values
}

func ReadFileAsArrayOfStringArrays(path string, delimiter string) [][]string {
	var values [][]string
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		values = append(values, strings.Split(scanner.Text(), delimiter))
	}
	_ = file.Close()
	return values
}
