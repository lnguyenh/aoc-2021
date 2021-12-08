package day08

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func doPart1(input [][][][]rune) int {
	counter := 0
	targetLengths := []int{2, 3, 4, 7}
	for _, entry := range input {
		for _, word := range entry[1] {
			wordLength := len(word)
			if utils.IntIn(wordLength, targetLengths) {
				counter++
			}
		}
	}
	return counter
}

func doPart2(input [][][][]rune) int {
	return 0
}

func massageInput(path string) [][][][]rune {
	var input [][][][]rune
	lines := utils.ReadFileAsStringSlice(path, "\n")
	for _, line := range lines {
		parts := utils.ParseStringAsStringList(line, "|")
		var lineParts [][][]rune
		for _, linePart := range parts {
			var partWords [][]rune
			words := utils.ParseStringAsStringList(linePart, " ")
			for _, word := range words {
				partWords = append(partWords, []rune(word))
			}
			lineParts = append(lineParts, partWords)
		}
		input = append(input, lineParts)
	}
	return input
}

func Run(path string) {
	input := massageInput(path)
	// fmt.Printf("input: %v\n", input)
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
