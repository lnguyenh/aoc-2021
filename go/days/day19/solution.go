package day19

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"sort"
	"strings"
)

func getBeaconKey(scannerId, beaconId int) string {
	return fmt.Sprintf("s%vb%v", scannerId, beaconId)
}

func getCoordinateKey(x, y, z int) string {
	return fmt.Sprintf("x%vx%vz%x", x, y, z)
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

func matchKey(a, b int) string {
	keys := []int{a, b}
	sort.Ints(keys)
	return fmt.Sprintf("%v", keys)
}

func doPart1(scanners []*aocScanner) int {
	matches := make([]scannerMatch, 0)
	existingMatches := make(map[string]bool)

	systemId := 0
	refScannerId := 0
	seen := []int{0}
	for {
		if len(seen) == len(scanners) {
			break
		}

		refScanner := scanners[refScannerId]

		for targetScannerId, targetScanner := range scanners {
			potentialMatchKey := matchKey(refScannerId, targetScannerId)
			if refScannerId == targetScannerId || existingMatches[potentialMatchKey] {
				continue
			}
			isAMatch, targetSystem, offset := refScanner.hasEnoughCommonPoints(targetScanner, systemId)
			if isAMatch {
				matches = append(matches, scannerMatch{
					refScannerId:    refScannerId,
					targetScannerId: targetScannerId,
					refSystem:       systemId,
					targetSystem:    targetSystem,
					offset:          offset,
				})
				systemId = targetSystem
				existingMatches[potentialMatchKey] = true
				seen = append(seen, targetScannerId)
				seen = utils.IntSliceToSet(seen)
			}
		}

	}

	points := make(map[string]bool)
	variableOffset := aocVector{x: 0, y: 0, z: 0}
	addPoints(points, scanners[0], 0, variableOffset)
	for _, match := range matches {
		variableOffset.x = variableOffset.x + match.offset.x
		variableOffset.y = variableOffset.y + match.offset.y
		variableOffset.z = variableOffset.z + match.offset.z
		addPoints(points, scanners[match.targetScannerId], match.targetSystem, variableOffset)
	}

	return 0
}

func doPart2(scanners []*aocScanner) int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	scanners := createScanners(input)
	scanners[0].print(1)

	fmt.Printf("scanners (%v): %v\n", len(scanners), scanners)
	answer1 := doPart1(scanners)
	answer2 := doPart2(scanners)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
