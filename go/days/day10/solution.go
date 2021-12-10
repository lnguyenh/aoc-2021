package day10

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"sort"
)

var opening = []rune{'(', '[', '{', '<'}

var closing = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

var scoring = map[rune]int{
	')': 1,
	'}': 3,
	']': 2,
	'>': 4,
}

func traverse(blob []rune) (bool, bool, rune, []rune) {
	var stack []rune
	for _, char := range blob {
		if utils.RuneInSlice(char, opening) {
			stack = append(stack, char)
		} else {
			if char == closing[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false, false, char, stack
			}
		}
	}
	if len(stack) != 0 {
		// incomplete
		return true, true, '0', stack
	} else {
		return true, false, '0', stack
	}

}

func doPart1(input [][]rune) int {
	errors := map[rune]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}
	for _, line := range input {
		isValid, _, char, _ := traverse(line)
		if !isValid {
			errors[char] += 1
		}
	}
	return errors[')']*3 + errors[']']*57 + errors['}']*1197 + errors['>']*25137
}

func generateMissingCharacters(stack []rune) []rune {
	missingChars := make([]rune, len(stack))
	for i, char := range stack {
		missingChars[i] = closing[char]
	}
	return utils.RuneRevertSlice(missingChars)
}

func doPart2(input [][]rune) int {
	var scores []int
	for _, line := range input {
		_, isIncomplete, _, stack := traverse(line)
		if isIncomplete {
			missingChars := generateMissingCharacters(stack)
			score := 0
			for _, char := range missingChars {
				score = score*5 + scoring[char]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfRuneSlices(path)
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
