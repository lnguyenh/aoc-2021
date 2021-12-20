package day20

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func run(image *aocImage) (int, int) {
	// Part 1
	image.applyOnce()
	image.applyOnce()
	answer1 := image.countOnes()

	// Part2, do it 48 more times
	for i := 2; i < 50; i++ {
		image.applyOnce()
	}
	answer2 := image.countOnes()
	return answer1, answer2
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	image := getImage(input)
	answer1, answer2 := run(image)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
