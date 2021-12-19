package day19

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

func getBeaconKey(scannerId, beaconId int) string {
	return fmt.Sprintf("s%vb%v", scannerId, beaconId)
}

func numCommonVectors(vectors1, vectors2 map[string]aocVector) int {
	count := 0
	for _, v1 := range vectors1 {
		for _, v2 := range vectors2 {
			if v1 == v2 {
				count++
			}
		}
	}
	return count
}

func textToCoordinates(blob string) aocCoordinates {
	var coordinates [3]int
	values := strings.Split(blob, ",")
	for i, value := range values {
		coordinates[i] = utils.StringToInt(value)
	}
	return aocCoordinates{x: coordinates[0], y: coordinates[1], z: coordinates[2]}
}

func createScanner(blob string) *aocScanner {
	lines := strings.Split(blob, "\n")
	line0 := strings.Split(lines[0], " ")

	// initialize
	scanner := aocScanner{
		id:              utils.StringToInt(line0[2]),
		originalBeacons: make(map[string]aocCoordinates),
		vectors:         make(map[string]*[24]map[string]aocVector),
	}
	for i, line := range lines[1:] {
		scanner.originalBeacons[getBeaconKey(scanner.id, i)] = textToCoordinates(line)
	}
	scanner.populateSystems()
	scanner.populateVectors()
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
	scanners[0].has12CommonPoints(scanners[1])
	fmt.Printf("scanners (%v): %v\n", len(scanners), scanners)
	answer1 := doPart1()
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
