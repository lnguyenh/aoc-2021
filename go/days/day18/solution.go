package day18

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func add(expression1, expression2 string) string {
	return strings.Join([]string{
		"[",
		expression1,
		",",
		expression2,
		"]"}, "")
}

func needsExplode(expression string) int {
	blob := []rune(expression)
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

func getNewValue(a, b rune) string {
	return fmt.Sprintf("%v", utils.RuneToInt(a)+utils.RuneToInt(b))
}

func insert(mainString, subString string, index int) string {
	newString := mainString[:index] + subString + mainString[index+1:]
	return newString
}

func explode(expression string, indexExplosion int) string {
	blob := []rune(expression)
	iLeft := indexExplosion + 1
	iRight := indexExplosion + 3
	iLeftBracket := indexExplosion
	iRightBracket := indexExplosion + 4
	var newLeftValue, newRightValue, leftPart, rightPart string

	// left merge
	leftPart = string(blob[:iLeftBracket])
	for i := iLeftBracket; i >= 0; i-- {
		if unicode.IsDigit(blob[i]) {
			newLeftValue = getNewValue(blob[i], blob[iLeft])
			leftPart = insert(leftPart, newLeftValue, i)
			break
		}
	}

	// right merge
	rightPart = string(blob[iRightBracket+1:])
	for i := iRightBracket; i < len(blob); i++ {
		if unicode.IsDigit(blob[i]) {
			newRightValue = getNewValue(blob[i], blob[iRight])
			insertionIndex := i - iRightBracket - 1
			rightPart = insert(rightPart, newRightValue, insertionIndex)
			break
		}
	}

	result := leftPart + "0" + rightPart
	return result
}

func split(expression string) string {
	blob := []rune(expression)
	for i := 0; i < len(blob)-1; i++ {
		potentialNumber, error := strconv.Atoi(string(blob[i : i+2]))
		if error == nil {
			up := utils.DigitToRune(int(math.Ceil(float64(potentialNumber) / 2)))
			down := utils.DigitToRune(int(math.Floor(float64(potentialNumber) / 2.0)))
			insert := []rune{'[', down, ',', up, ']'}
			leftPart := blob[:i]
			rightPart := []rune(string(blob[i+2:]))
			processed := append(leftPart, insert...)
			processed = append(processed, rightPart...)
			return string(processed)
		}
	}
	return string(blob)
}

func doTests() {
	expression := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	fmt.Println(expression)
	indexExplosion := needsExplode(expression)
	exploded := explode(expression, indexExplosion)
	fmt.Printf("%v\n", exploded)

	expression = "[[[[0,7],4],[15,[0,13]]],[1,1]]"
	fmt.Println(expression)
	splitted := split(expression)
	fmt.Printf("%v\n", splitted)
}

func doPart1(expressions []string) int {
	expression := expressions[0]
	for _, nextExpression := range expressions[1:] {
		expression = add(expression, nextExpression)
		for {
			indexExplosion := needsExplode(expression)
			if indexExplosion > 0 {
				expression = explode(expression, indexExplosion)
				expression = split(expression)
				expression = split(expression)
			} else {
				break
			}
		}
	}
	fmt.Println(expression)
	return 0
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	fmt.Printf("input: %v\n", input)
	answer1 := doPart1(input)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
