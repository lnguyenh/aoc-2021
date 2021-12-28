package day22

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
)

func doPart2(instructions []string) int {
	r, _ := regexp.Compile("^(\\w+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)$")
	parsedInstructions := make([]aocInstruction, 0, len(instructions))
	for _, instruction := range instructions {
		groups := r.FindStringSubmatch(instruction)
		parsedInstructions = append(parsedInstructions, aocInstruction{
			isFull:           groups[1] == "on",
			cuboidBoundaries: utils.StringSliceToIntSlice(groups[2:])})
	}
	space := createSpace(parsedInstructions)
	space.buildAxes()
	space.simplify()
	space.applyInstructions()
	return space.getVolume()
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	answer1 := doPart1(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := doPart2(input)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
