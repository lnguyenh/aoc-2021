package day18

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"math"
	"regexp"
	"unicode"
)

func add(expression1, expression2 string) string {
	return "[" + expression1 + "," + expression2 + "]"
}

func needsExplode(expression string) int {
	blob := []rune(expression)
	var stack []rune
	r, _ := regexp.Compile("^\\[\\d+\\,\\d+\\]")

	for i := 0; i < len(blob); i++ {
		char1 := blob[i]
		if char1 == '[' {
			stack = append(stack, char1)
		} else if char1 == ']' {
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) > 4 {
				match := r.MatchString(string(blob[i-1:]))
				if match {
					return i - 1
				}
			}
		}
	}

	return -1
}

func getMagnitude(expression string) int {
	blob := []rune(expression)
	var stack []rune
	var values []int

	for i := 0; i < len(blob); i++ {
		char := blob[i]
		if char == '[' {
			stack = append(stack, char)
		} else if char == ']' {
			newValue := 3*values[len(values)-2] + 2*values[len(values)-1]
			values = values[:len(values)-2]
			values = append(values, newValue)
			stack = stack[:len(stack)-1]
		} else if unicode.IsDigit(char) {
			values = append(values, utils.RuneToInt(char))
		}
	}
	return values[0]
}

func getLeftInt(blob []rune) (int, int, int) {
	for i := len(blob) - 1; i > 0; i-- {
		if unicode.IsDigit(blob[i]) {
			stopIndex := i
			for j := i - 1; j > 0; j-- {
				if !unicode.IsDigit(blob[j]) {
					startIndex := j + 1
					value := utils.StringToInt(string(blob[startIndex : stopIndex+1]))
					return value, startIndex, stopIndex
				}
			}
		}
	}
	return -1, -1, -1
}

func getRightInt(blob []rune) (int, int, int) {
	for i := 0; i < len(blob); i++ {
		if unicode.IsDigit(blob[i]) {
			startIndex := i
			for j := i + 1; j < len(blob); j++ {
				if !unicode.IsDigit(blob[j]) {
					stopIndex := j - 1
					value := utils.StringToInt(string(blob[startIndex : stopIndex+1]))
					return value, startIndex, stopIndex
				}
			}
		}
	}
	return -1, -1, -1
}

func insert(mainString, subString string, startIndex, stopIndex int) string {
	newString := mainString[:startIndex] + subString + mainString[stopIndex+1:]
	return newString
}

func explode(expression string, indexExplosion int) string {
	// indexExplosion is index of left bracket
	blob := []rune(expression)
	var newLeftValue, newRightValue, leftPart, rightPart string

	r, _ := regexp.Compile("^\\[(\\d+)\\,(\\d+)\\]")
	matches := r.FindStringSubmatch(string(blob[indexExplosion:]))

	// left merge
	leftPart = string(blob[:indexExplosion])
	leftInt, leftStartIndex, leftStopIndex := getLeftInt([]rune(leftPart))
	if leftInt >= 0 {
		newLeftValue = fmt.Sprintf("%v", leftInt+utils.StringToInt(matches[1]))
		leftPart = insert(leftPart, newLeftValue, leftStartIndex, leftStopIndex)
	}

	// right merge
	rightPart = string(blob[indexExplosion+len(matches[0]):])
	rightInt, rightStartIndex, rightStopIndex := getRightInt([]rune(rightPart))
	if rightInt >= 0 {
		newRightValue = fmt.Sprintf("%v", rightInt+utils.StringToInt(matches[2]))
		rightPart = insert(rightPart, newRightValue, rightStartIndex, rightStopIndex)
	}

	result := leftPart + "0" + rightPart
	return result
}

func split(expression string) string {
	blob := []rune(expression)
	for i := 0; i < len(blob)-1; i++ {

		r, _ := regexp.Compile("^(\\d+)")
		if r.MatchString(string(blob[i:])) {
			number, startIndex, stopIndex := getRightInt(blob[i:])
			if number > 9 {
				up := int(math.Ceil(float64(number) / 2))
				down := int(math.Floor(float64(number) / 2))
				insertionString := fmt.Sprintf("[%v,%v]", down, up)
				return insert(expression, insertionString, i+startIndex, i+stopIndex)
			}
		}
	}
	return string(blob)
}

func doTests() {
	expression := "[[[[[[[[[11,[0,[15,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]"
	expression = reduce(expression)
	fmt.Println(getLeftInt([]rune("wewe23")))
	fmt.Println(getRightInt([]rune("a25we")))
}

func reduce(expressionToReduce string) string {
	expression := expressionToReduce
	for {
		originalLength := len(expression)

		indexExplosion := needsExplode(expression)
		if indexExplosion > 0 {
			// fmt.Printf("%v needs explode at %v\n", expression, indexExplosion)
			expression = explode(expression, indexExplosion)
			// fmt.Printf("%v after explode\n", expression)
		} else {
			expression = split(expression)
			// fmt.Printf("%v after split\n", expression)
		}

		if originalLength == len(expression) {
			break
		}
	}
	return expression
}

func doPart1(expressions []string) int {
	expression := expressions[0]
	for _, nextExpression := range expressions[1:] {
		expression = add(expression, nextExpression)
		expression = reduce(expression)

	}
	return getMagnitude(expression)
}

func doPart2(expressions []string) int {
	asMap := make(map[int]string)
	magnitudes := make([]int, 0)
	for i, expression := range expressions {
		asMap[i] = expression
	}
	for key1, expression1 := range asMap {
		for key2, expression2 := range asMap {
			if key1 != key2 {
				magnitudes = append(magnitudes, getMagnitude(reduce(add(expression1, expression2))))
			}
		}
	}
	return utils.MaxSlice(magnitudes)

}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
