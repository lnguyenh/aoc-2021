package day19

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

func getBeaconKey(scannerId, beaconId int) string {
	return fmt.Sprintf("s%vb%v", scannerId, beaconId)
}

func textToCoordinates(blob string) [3]int {
	var coordinates [3]int
	values := strings.Split(blob, ",")
	for i, value := range values {
		coordinates[i] = utils.StringToInt(value)
	}
	return coordinates
}

func createScanner(blob string) *aocScanner {
	lines := strings.Split(blob, "\n")
	line0 := strings.Split(lines[0], " ")

	// initialize
	scanner := aocScanner{
		id:              utils.StringToInt(line0[2]),
		originalBeacons: make(map[string][3]int),
	}
	for i, line := range lines[1:] {
		scanner.originalBeacons[getBeaconKey(scanner.id, i)] = textToCoordinates(line)
	}
	return &scanner
}

func createScanners(blobs []string) []*aocScanner {
	scanners := make([]*aocScanner, 0, len(blobs))
	for _, blob := range blobs {
		scanners = append(scanners, createScanner(blob))
	}
	return scanners
}

func doPart1() int {
	return 0
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	scanners := createScanners(input)
	fmt.Printf("scanners (%v): %v\n", len(scanners), scanners)
	answer1 := doPart1()
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
