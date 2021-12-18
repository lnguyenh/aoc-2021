package day18

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"unicode"
)

func needsExplode(blob []rune) int {
	var stack []rune
	for i, char := range blob {
		if char == '[' {
			stack = append(stack, char)
		} else if char == ']' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 5 {
			return i
		}
	}
	return -1
}

func reduce(blob []rune, indexExplosion int) string {
	iLeft := indexExplosion + 1
	iRight := indexExplosion + 3
	iLeftBracket := indexExplosion
	iRightBracket := indexExplosion + 4

	// left merge
	for i := iLeftBracket; i >= 0; i-- {
		if unicode.IsDigit(blob[i]) {
			newValue := utils.DigitToRune(utils.RuneToInt(blob[i]) + utils.RuneToInt(blob[iLeft]))
			blob[i] = newValue
			break
		}
	}

	// right merge
	for i := iRightBracket; i < len(blob); i++ {
		if unicode.IsDigit(blob[i]) {
			newValue := utils.DigitToRune(utils.RuneToInt(blob[i]) + utils.RuneToInt(blob[iRight]))
			blob[i] = newValue
			break
		}
	}

	leftPart := blob[:iLeftBracket]
	rightPart := blob[iRightBracket+1:]
	exploded := append(leftPart, '0')
	exploded = append(exploded, rightPart...)
	return string(exploded)
}

func doPart1() int {
	expression := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	asRunes := []rune(expression)
	fmt.Println(expression)
	indexExplosion := needsExplode([]rune(expression))
	reduced := reduce(asRunes, indexExplosion)
	fmt.Printf("%v\n", reduced)

	return 0
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	fmt.Printf("input: %v\n", input)
	answer1 := doPart1()
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
