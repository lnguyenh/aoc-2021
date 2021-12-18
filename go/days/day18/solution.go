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
	for i := 0; i < len(blob)-3; i++ {
		char1, char2, char3 := blob[i], blob[i+1], blob[i+2]
		if char1 == '[' {
			stack = append(stack, char1)
		} else if char1 == ']' {
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) >= 5 && unicode.IsDigit(char1) && char2 == ',' && unicode.IsDigit(char3) {
				fmt.Printf("%v\n", string(char1))
				return i - 1
			}
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
		potentialNumber, e := strconv.Atoi(string(blob[i : i+2]))
		if e == nil {
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
	for {
		originalLength := len(expression)
		indexExplosion := needsExplode(expression)
		if indexExplosion > 0 {
			fmt.Printf("%v needs explode at %v\n", expression, indexExplosion)
			expression = explode(expression, indexExplosion)
			fmt.Printf("%v after explode\n", expression)
		} else {
			expression = split(expression)
			fmt.Printf("%v after split\n", expression)
		}
		if originalLength == len(expression) {
			break
		}
	}
}

func doPart1(expressions []string) int {
	expression := expressions[0]
	fmt.Println("----------------------------------------------")
	for i, nextExpression := range expressions[1:] {
		fmt.Println("----------------------------------------------")
		fmt.Printf("#%v\n", i)
		fmt.Printf("%v expression \n", expression)
		fmt.Printf("%v next expression\n", nextExpression)
		expression = add(expression, nextExpression)
		fmt.Printf("%v after addition\n", expression)
		for {
			indexExplosion := needsExplode(expression)
			if indexExplosion > 0 {
				fmt.Printf("%v needs explode at %v\n", expression, indexExplosion)
				expression = explode(expression, indexExplosion)
				fmt.Printf("%v after explode\n", expression)
				expression = split(expression)
				fmt.Printf("%v after split\n", expression)
				expression = split(expression)
				fmt.Printf("%v after split\n", expression)
			} else {
				break
			}
		}
	}
	fmt.Println(expression)
	return 0
}

func doPart2(expressions []string) int {
	return 0
}

var onlyTests = false

func Run(path string) {
	if !onlyTests {
		input := utils.ReadFileAsStringSlice(path, "\n")
		answer1 := doPart1(input)
		answer2 := doPart2(input)
		fmt.Printf("Part 1 answer: %v\n", answer1)
		fmt.Printf("Part 2 answer: %v\n", answer2)
	} else {
		doTests()
	}
}
