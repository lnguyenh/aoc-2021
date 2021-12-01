package filereader

import (
	"bufio"
	"os"
	"strconv"
)

func ReadAsIntArray(path string) []int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var values []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}
	file.Close()
	return values
}

func ReadAsStringArray(path string) []string {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	file.Close()
	return values
}
