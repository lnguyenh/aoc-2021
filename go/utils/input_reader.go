package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsIntSlice(path string) []int {
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

func ReadFileAsString(path string) string {
	var lines []string
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_ = file.Close()
	return strings.Join(lines, "\n")
}

func ReadFileAsStringSlice(path string, delimiter string) []string {
	return strings.Split(ReadFileAsString(path), delimiter)
}

func ReadFileAsSliceOfStringSlices(path string, delimiter string) [][]string {
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

func ReadFileAsSliceOfRuneSlices(path string) [][]rune {
	var values [][]rune
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		values = append(values, []rune(scanner.Text()))
	}
	_ = file.Close()
	return values
}
