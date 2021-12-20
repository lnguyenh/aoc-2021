package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

type aocImage struct {
	points        map[string]rune
	algorithm     []rune
	infinityValue rune
	minI          int
	maxI          int
	minJ          int
	maxJ          int
}

func (image *aocImage) applyOnce() {
	result := make(map[string]rune)

	// Convert all points, and one row/column extra all around
	image.minI--
	image.minJ--
	image.maxI++
	image.maxJ++
	for i := image.minI; i <= image.maxI; i++ {
		for j := image.minJ; j <= image.maxJ; j++ {
			key := getKey(i, j)
			result[key] = image.algorithm[getAlgorithmIndex(image.points, key, image.infinityValue)]
		}
	}

	//printImage(result)
	image.infinityValue = result[getKey(image.minI, image.minJ)]
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
	matches := strings.Split(key, "Z")
	return utils.StringToInt(matches[0]), utils.StringToInt(matches[1])
}
