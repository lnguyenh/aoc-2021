package day08

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strconv"
)

// These are the patterns that all lines should have
var referencePatterns = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

// Conversion table from pattern to number as a rune
var wordToNumberRune = map[string]rune{
	"abcefg":  '0',
	"cf":      '1',
	"acdeg":   '2',
	"acdfg":   '3',
	"bcdf":    '4',
	"abdfg":   '5',
	"abdefg":  '6',
	"acf":     '7',
	"abcdefg": '8',
	"abcdfg":  '9',
}

func generateCombinations(signalPatterns [][]rune) [][]rune {
	var a rune
	bOrD := make([]rune, 0, 2)
	cOrF := make([]rune, 0, 2)
	eOrG := make([]rune, 0, 2)

	for _, word := range signalPatterns {
		if len(word) == 2 {
			cOrF = word
		}
	}

	for _, word := range signalPatterns {
		if len(word) == 3 {
			for _, letter := range word {
				if !utils.RuneInSlice(letter, cOrF) {
					a = letter
				}
			}
		}
	}

	for _, word := range signalPatterns {
		if len(word) == 4 {
			for _, letter := range word {
				if !utils.RuneInSlice(letter, cOrF) {
					bOrD = append(bOrD, letter)
				}
			}
		}
	}

	for _, word := range signalPatterns {
		if len(word) == 7 {
			for _, letter := range word {
				if !utils.RuneInSlice(letter, cOrF) && !utils.RuneInSlice(letter, bOrD) && letter != a {
					eOrG = append(eOrG, letter)
				}
			}
		}
	}

	// each combination is a candidate translation for "abcdefg"
	var combinations [][]rune
	combinations = append(combinations, []rune{a, bOrD[0], cOrF[0], bOrD[1], eOrG[0], cOrF[1], eOrG[1]})
	combinations = append(combinations, []rune{a, bOrD[0], cOrF[0], bOrD[1], eOrG[1], cOrF[1], eOrG[0]})
	combinations = append(combinations, []rune{a, bOrD[1], cOrF[0], bOrD[0], eOrG[0], cOrF[1], eOrG[1]})
	combinations = append(combinations, []rune{a, bOrD[1], cOrF[0], bOrD[0], eOrG[1], cOrF[1], eOrG[0]})
	combinations = append(combinations, []rune{a, bOrD[0], cOrF[1], bOrD[1], eOrG[0], cOrF[0], eOrG[1]})
	combinations = append(combinations, []rune{a, bOrD[0], cOrF[1], bOrD[1], eOrG[1], cOrF[0], eOrG[0]})
	combinations = append(combinations, []rune{a, bOrD[1], cOrF[1], bOrD[0], eOrG[0], cOrF[0], eOrG[1]})
	combinations = append(combinations, []rune{a, bOrD[1], cOrF[1], bOrD[0], eOrG[1], cOrF[0], eOrG[0]})

	return combinations
}

func translate(table map[rune]rune, word []rune) []rune {
	var translatedWord []rune
	for _, letter := range word {
		translatedLetter := table[letter]
		translatedWord = append(translatedWord, translatedLetter)
	}
	return utils.SortRuneSlice(translatedWord)
}

func isValid(table map[rune]rune, signalPatterns [][]rune) bool {
	for _, pattern := range signalPatterns {
		translatedPattern := translate(table, pattern)
		if !utils.StringInSlice(string(translatedPattern), referencePatterns) {
			return false
		}
	}
	return true
}

func applyCombination(table map[rune]rune, words [][]rune) int {
	var numberAsRunes []rune
	for _, word := range words {
		numberAsRunes = append(numberAsRunes, wordToNumberRune[string(translate(table, word))])
	}
	number, _ := strconv.Atoi(string(numberAsRunes))
	return number
}

func generateTranslationTable(combination []rune) map[rune]rune {
	table := make(map[rune]rune)
	for i, letter := range []rune("abcdefg") {
		table[combination[i]] = letter
	}
	return table
}

func doPart2(input [][][][]rune) int {
	counter := 0
	for _, entry := range input {
		combinations := generateCombinations(entry[0])
		for _, combination := range combinations {
			translationTable := generateTranslationTable(combination)
			if isValid(translationTable, entry[0]) {
				counter += applyCombination(translationTable, entry[1])
			}
		}

	}
	return counter
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

func doPart1(input [][][][]rune) int {
	counter := 0
	targetLengths := []int{2, 3, 4, 7}
	for _, entry := range input {
		for _, word := range entry[1] {
			wordLength := len(word)
			if utils.IntInSlice(wordLength, targetLengths) {
				counter++
			}
		}
	}
	return counter
}

func Run(path string) {
	input := massageInput(path)
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
