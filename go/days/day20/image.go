package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
)

type aocImage struct {
	points        map[string]rune
	algorithm     []rune
	infinityValue rune
}

func (image *aocImage) applyOnce() {
	result := make(map[string]rune)

	// vanilla run on base image
	for key := range image.points {
		result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
	}

	// continue until we have the same character 3 character wide around the image
	for {
		canStop, newInfinityValue := isReadyForInfinite(result)
		if canStop {
			image.infinityValue = newInfinityValue
			break
		}

		// continue once around
		minI, maxI, minJ, maxJ := getMinMax(result)
		for j := minJ - 1; j <= maxJ+1; j++ {
			key := getKey(minI-1, j)
			result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
		}
		for j := minJ - 1; j <= maxJ+1; j++ {
			key := getKey(maxI+1, j)
			result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
		}
		for i := minI - 1; i <= maxI+1; i++ {
			key := getKey(i, minJ-1)
			result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
		}
		for i := minI - 1; i <= maxI+1; i++ {
			key := getKey(i, maxJ+1)
			result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
		}

	}

	image.points = result
}

func printImage(points map[string]rune) {
	minI, maxI, minJ, maxJ := getMinMax(points)
	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if points[getKey(i, j)] == '0' {
				fmt.Printf(". ")
			} else if points[getKey(i, j)] == '1' {
				fmt.Printf("# ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func (image *aocImage) countOnes() int {
	counter := 0
	for _, val := range image.points {
		if val == '1' {
			counter++
		}
	}
	return counter
}

func isReadyForInfinite(points map[string]rune) (bool, rune) {
	minI, maxI, minJ, maxJ := getMinMax(points)
	margin := 3
	refChar := points[getKey(minI, minJ)]
	for i := minI; i <= maxJ; i++ {
		for j := minJ; j <= maxJ; j++ {
			if minI+margin < i && i < maxI-margin &&
				minJ+margin < j && j < maxJ-margin {
				continue
			}
			testChar := points[getKey(i, j)]
			if testChar != refChar {
				return false, refChar
			}
		}
	}
	return true, refChar
}

func getVal(image map[string]rune, i, j int, fallback rune) rune {
	key := getKey(i, j)
	char, ok := image[key]
	if ok {
		return char
	} else {
		return fallback
	}
}

func getAlgorithmIndex(points map[string]rune, key string, infinityValue rune) int {
	i, j := keyToIJ(key)
	code := string([]rune{
		getVal(points, i-1, j-1, infinityValue),
		getVal(points, i-1, j, infinityValue),
		getVal(points, i-1, j+1, infinityValue),
		getVal(points, i, j-1, infinityValue),
		getVal(points, i, j, infinityValue),
		getVal(points, i, j+1, infinityValue),
		getVal(points, i+1, j-1, infinityValue),
		getVal(points, i+1, j, infinityValue),
		getVal(points, i+1, j+1, infinityValue),
	})
	return utils.BitsToInt(code)
}

func getMinMax(image map[string]rune) (int, int, int, int) {
	sliceI := make([]int, 0, len(image))
	sliceJ := make([]int, 0, len(image))
	for key := range image {
		i, j := keyToIJ(key)
		sliceI = append(sliceI, i)
		sliceJ = append(sliceJ, j)
	}
	minI, maxI := utils.MinSlice(sliceI), utils.MaxSlice(sliceI)
	minJ, maxJ := utils.MinSlice(sliceJ), utils.MaxSlice(sliceJ)
	return minI, maxI, minJ, maxJ
}

func getKey(i, j int) string {
	return fmt.Sprintf("%vZ%v", i, j)
}

func keyToIJ(key string) (int, int) {
	r, _ := regexp.Compile("^(-?\\d+)Z(-?\\d+)")
	matches := r.FindStringSubmatch(key)
	return utils.StringToInt(matches[1]), utils.StringToInt(matches[2])
}
