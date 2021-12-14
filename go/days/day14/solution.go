package day14

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
	"unicode"
)

func getAlphabet() []rune {
	alphabet := make([]rune, 0, 25)
	for r := 'a'; r < 'z'; r++ {
		alphabet = append(alphabet, unicode.ToUpper(r))
	}
	return alphabet
}

func getCountsBetter(segmentCounts map[string]int, endLetter string) (map[string]int, []int) {
	counts := make(map[string]int)
	countsAsSlice := make([]int, 0, 3)
	for _, letter := range getAlphabet() {
		counts[string(letter)] = 0
	}
	for segment, count := range segmentCounts {
		counts[string(segment[0])] += count
	}
	counts[endLetter] += 1
	for _, count := range counts {
		if count > 0 {
			countsAsSlice = append(countsAsSlice, count)
		}
	}
	return counts, countsAsSlice
}

func getConversionsBetter(text string) map[string][2]string {
	conversions := make(map[string][2]string)
	for _, line := range strings.Split(text, "\n") {
		elements := strings.Split(line, " -> ")
		initialSegment := elements[0]
		resultingSegments := [2]string{
			elements[0][0:1] + elements[1],
			elements[1] + elements[0][1:2],
		}
		conversions[initialSegment] = resultingSegments
	}
	return conversions
}

func applyConversionBetter(initial map[string]int, conversions map[string][2]string) map[string]int {
	var output = make(map[string]int)
	for segment, count := range initial {
		segments := conversions[segment]
		output[segments[0]] += count
		output[segments[1]] += count
	}
	return output
}

func getInitialCounts(text string, conversions map[string][2]string) map[string]int {
	var initialCounts = make(map[string]int)
	for segment := range conversions {
		initialCounts[segment] = 0
	}
	for i := 0; i < len(text)-1; i++ {
		initialCounts[text[i:i+2]] += 1
	}
	return initialCounts
}

func doPart1(initial map[string]int, conversions map[string][2]string, endLetter string) int {
	var transformed = initial
	for i := 0; i < 10; i++ {
		transformed = applyConversionBetter(transformed, conversions)
	}
	_, counts := getCountsBetter(transformed, endLetter)
	return utils.MaxSlice(counts) - utils.MinSlice(counts)

}

func doPart2(initial map[string]int, conversions map[string][2]string, endLetter string) int {
	var transformed = initial
	for i := 0; i < 40; i++ {
		transformed = applyConversionBetter(transformed, conversions)
	}
	_, counts := getCountsBetter(transformed, endLetter)
	return utils.MaxSlice(counts) - utils.MinSlice(counts)
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	conversions := getConversionsBetter(input[1])
	initial := getInitialCounts(input[0], conversions)
	endLetter := string(input[0][len(input[0])-1])
	answer1 := doPart1(initial, conversions, endLetter)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := doPart2(initial, conversions, endLetter)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
