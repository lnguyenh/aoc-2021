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

func run(scanners []*aocScanner) (int, int) {
	matches := make([]scannerMatch, 0)
	existingMatches := make(map[string]bool)

	systemId := 0
	for refScannerId, refScanner := range scanners {
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
			}
		}
	}

	standards := make(map[int]aocStandard)
	standards[0] = aocStandard{offset: aocVector{}, system: 0}
	added := make(map[int]bool)
	added[0] = true
	for {
		if len(standards) == len(scanners) {
			break
		}

		for scannerId, scanner := range scanners {
			if added[scannerId] {
				continue
			}
			for _, match := range matches {
				if added[match.refScannerId] && match.targetScannerId == scannerId {
					// refScanner in the match is linked to scanner 0
					standard := standards[match.refScannerId]
					knownScanner := scanners[match.refScannerId]
					_, newSystem, newOffset := knownScanner.hasEnoughCommonPoints(scanner, standard.system)
					standards[scannerId] = aocStandard{
						offset: aocVector{
							x: newOffset.x + standard.offset.x,
							y: newOffset.y + standard.offset.y,
							z: newOffset.z + standard.offset.z,
						},
						system: newSystem,
					}
					added[scannerId] = true
				} else if added[match.targetScannerId] && match.refScannerId == scannerId {
					// targetScanner in the match is linked to scanner 0
					standard := standards[match.targetScannerId]
					knownScanner := scanners[match.targetScannerId]
					_, newSystem, newOffset := knownScanner.hasEnoughCommonPoints(scanner, standard.system)
					standards[scannerId] = aocStandard{
						offset: aocVector{
							x: newOffset.x + standard.offset.x,
							y: newOffset.y + standard.offset.y,
							z: newOffset.z + standard.offset.z,
						},
						system: newSystem,
					}
					added[scannerId] = true
				}
			}
		}

	}

	points := make(map[string]bool)
	for scannerId, standard := range standards {
		addPoints(points, scanners[scannerId], standard.system, standard.offset)
	}

	manhattans := make([]int, 0, len(standards))
	for id1, standard1 := range standards {
		for id2, standard2 := range standards {
			if id1 != id2 {
				manhattans = append(
					manhattans,
					utils.IntAbs(standard1.offset.x-standard2.offset.x)+
						utils.IntAbs(standard1.offset.y-standard2.offset.y)+
						utils.IntAbs(standard1.offset.z-standard2.offset.z))
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
