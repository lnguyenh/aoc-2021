package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type aocImage struct {
	points        map[string]rune
	algorithm     []rune
	infinityValue rune
}

func (image *aocImage) applyOnce() map[string]rune {
	result := make(map[string]rune)
	for key := range image.points {
		i, j := keyToIJ(key)
		code := string([]rune{
			getVal(image.points, i-1, j-1, image.infinityValue),
			getVal(image.points, i-1, j, image.infinityValue),
			getVal(image.points, i-1, j+1, image.infinityValue),
			getVal(image.points, i, j-1, image.infinityValue),
			getVal(image.points, i, j, image.infinityValue),
			getVal(image.points, i, j+1, image.infinityValue),
			getVal(image.points, i+1, j-1, image.infinityValue),
			getVal(image.points, i+1, j, image.infinityValue),
			getVal(image.points, i+1, j+1, image.infinityValue),
		})
		index := utils.BitsToInt(code)
		result[key] = image.algorithm[index]
	}
	pad(result)
	return result
}

func (image *aocImage) print() {
	minI, maxI, minJ, maxJ := getMinMax(image.points)
	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if image.points[getKey(i, j)] == '0' {
				fmt.Printf(". ")
			} else if image.points[getKey(i, j)] == '1' {
				fmt.Printf("# ")
			}
		}
		fmt.Printf("\n")
	}
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

func pad(image map[string]rune) {
	minI, maxI, minJ, maxJ := getMinMax(image)
	pad := 0
	for i := minI - pad; i < minI; i++ {
		for j := minJ - pad; j <= maxJ+pad; j++ {
			image[getKey(i, j)] = '0'
		}
	}
	for i := minI; i <= maxI; i++ {
		for j := minJ - pad; j < minJ; j++ {
			image[getKey(i, j)] = '0'
		}
	}
	for i := minI; i <= maxI; i++ {
		for j := maxJ + 1; j <= maxJ+pad; j++ {
			image[getKey(i, j)] = '0'
		}
	}
	for i := maxI + 1; i <= maxI+pad; i++ {
		for j := minJ - pad; j <= maxJ+pad; j++ {
			image[getKey(i, j)] = '0'
		}
	}
}
