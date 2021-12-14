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

func getCounts(text string) (map[rune]int, []int) {
	counts := make(map[rune]int)
	countsAsSlice := make([]int, 0, 3)
	for _, letter := range getAlphabet() {
		counter := 0
		for _, char := range []rune(text) {
			if char == letter {
				counter++
			}
		}
		if counter > 0 {
			counts[letter] = counter
			countsAsSlice = append(countsAsSlice, counter)
		}
	}
	return counts, countsAsSlice
}

func applyConversion(initial string, conversions map[string][]rune) string {
	var output = make([]rune, 0, 2*len(initial)+1)
	output = append(output, rune(initial[0]))
	for i := 0; i < len(initial)-1; i++ {
		key := initial[i : i+2]
		output = append(output, conversions[key]...)
	}
	return string(output)
}

func doPart1(initial string, conversions map[string][]rune) int {
	var transformed = initial

	for i := 0; i < 10; i++ {
		transformed = applyConversion(transformed, conversions)
		// fmt.Println(transformed)
	}
	_, counts := getCounts(transformed)
	return utils.MaxSlice(counts) - utils.MinSlice(counts)

}

func doPart2(initial string, conversions map[string][]rune) int {
	var transformed = initial

	for i := 0; i < 40; i++ {
		transformed = applyConversion(transformed, conversions)
		// fmt.Println(transformed)
	}
	_, counts := getCounts(transformed)
	return utils.MaxSlice(counts) - utils.MinSlice(counts)
}

func getConversions(text string) map[string][]rune {
	conversions := make(map[string][]rune)
	for _, line := range strings.Split(text, "\n") {
		elements := strings.Split(line, " -> ")
		conversions[elements[0]] = []rune(elements[1] + string(elements[0][1]))
	}
	return conversions
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n\n")
	// fmt.Printf("input: -%v-\n", input[1])
	// fmt.Printf("input: -%v-\n", input[0])
	initial := input[0]
	conversions := getConversions(input[1])
	// answer1 := doPart1(initial, conversions)
	// fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := doPart2(initial, conversions)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
