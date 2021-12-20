package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func doPart1(image *aocImage) int {
	image.applyOnce()
	image.applyOnce()
	return image.countOnes()
}

func doPart2(image *aocImage) int {
	for i := 0; i < 50; i++ {
		image.applyOnce()
	}
	return image.countOnes()
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")

	image := &aocImage{
		algorithm:     getAlgorithm(input[0]),
		points:        getImage(input[1]),
		infinityValue: '0',
	}
	image2 := &aocImage{
		algorithm:     getAlgorithm(input[0]),
		points:        getImage(input[1]),
		infinityValue: '0',
	}
	answer1 := doPart1(image)
	answer2 := doPart2(image2)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
