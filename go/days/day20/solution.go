package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
)

func getKey(i, j int) string {
	return fmt.Sprintf("%vZ%v", i, j)
}

func keyToIJ(key string) (int, int) {
	r, _ := regexp.Compile("^(-?\\d+)Z(-?\\d+)")
	matches := r.FindStringSubmatch(key)
	return utils.StringToInt(matches[1]), utils.StringToInt(matches[2])
}

func doPart1(image *aocImage) int {
	image.print()
	image.applyOnce()
	return image.countOnes()
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")

	image := &aocImage{
		algorithm:     getAlgorithm(input[0]),
		points:        getImage(input[1]),
		infinityValue: '0',
	}
	answer1 := doPart1(image)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
