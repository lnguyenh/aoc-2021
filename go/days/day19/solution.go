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

func findMatches(scanners []*aocScanner) []scannerMatch {
	matches := make([]scannerMatch, 0)
	existingMatches := make(map[string]bool)
	systemId := 0
	for refScannerId, refScanner := range scanners {
		for targetScannerId, targetScanner := range scanners {
			potentialMatchKey := matchKey(refScannerId, targetScannerId)
			if refScannerId == targetScannerId || existingMatches[potentialMatchKey] {
				continue
			}
			isAMatch, targetSystem, offset := refScanner.findNeighbouringData(targetScanner, systemId)
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
			}
		}
	}
	return matches
}

func findNormalizations(scanners []*aocScanner, matches []scannerMatch) map[int]aocNormalization {
	normalizations := make(map[int]aocNormalization)

	// Initialize with the first scanner
	normalizations[0] = aocNormalization{offset: aocVector{}, system: 0}
	added := make(map[int]bool)
	added[0] = true

	for {
		if len(normalizations) == len(scanners) {
			break
		}

		for scannerId, scanner := range scanners {
			if added[scannerId] {
				// Already normalized
				continue
			}
			for _, match := range matches {
				if added[match.refScannerId] && match.targetScannerId == scannerId {
					// refScanner in the match is normalized to scanner 0
					knownNormalization := normalizations[match.refScannerId]
					knownScanner := scanners[match.refScannerId]
					_, newSystem, newOffset := knownScanner.findNeighbouringData(scanner, knownNormalization.system)
					normalizations[scannerId] = aocNormalization{
						offset: aocVector{
							x: newOffset.x + knownNormalization.offset.x,
							y: newOffset.y + knownNormalization.offset.y,
							z: newOffset.z + knownNormalization.offset.z,
						},
						system: newSystem,
					}
					added[scannerId] = true
				} else if added[match.targetScannerId] && match.refScannerId == scannerId {
					// targetScanner in the match is normalized to scanner 0
					knownNormalization := normalizations[match.targetScannerId]
					knownScanner := scanners[match.targetScannerId]
					_, newSystem, newOffset := knownScanner.findNeighbouringData(scanner, knownNormalization.system)
					normalizations[scannerId] = aocNormalization{
						offset: aocVector{
							x: newOffset.x + knownNormalization.offset.x,
							y: newOffset.y + knownNormalization.offset.y,
							z: newOffset.z + knownNormalization.offset.z,
						},
						system: newSystem,
					}
					added[scannerId] = true
				}
			}
		}
	}
	return normalizations
}

func run(scanners []*aocScanner) (int, int) {
	matches := findMatches(scanners)
	normalizations := findNormalizations(scanners, matches)

	// Part 1: Build a map with all points
	points := make(map[string]bool)
	for scannerId, normalization := range normalizations {
		addPoints(points, scanners[scannerId], normalization.system, normalization.offset)
	}

	// Part 2: Manhattan
	manhattans := make([]int, 0, len(normalizations))
	for id1, normalization1 := range normalizations {
		for id2, normalization2 := range normalizations {
			if id1 != id2 {
				manhattans = append(
					manhattans,
					utils.IntAbs(normalization1.offset.x-normalization2.offset.x)+
						utils.IntAbs(normalization1.offset.y-normalization2.offset.y)+
						utils.IntAbs(normalization1.offset.z-normalization2.offset.z))
			}
		}
	}

	return len(points), utils.MaxSlice(manhattans)
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	scanners := createScanners(input)
	answer1, answer2 := run(scanners)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
