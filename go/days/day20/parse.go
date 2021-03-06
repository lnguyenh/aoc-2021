package day20

import (
	"fmt"
	"strings"
)

func getAlgorithm(blob string) []rune {
	result := make([]rune, 0, len(blob))
	for _, char := range []rune(blob) {
		var c rune
		switch char {
		case '.':
			c = '0'
		case '#':
			c = '1'
		default:
			fmt.Printf("ERROR ")
		}
		result = append(result, c)
	}
	return result
}

func getPoints(blob string) map[string]rune {
	result := make(map[string]rune)
	for i, line := range strings.Split(blob, "\n") {
		for j, char := range []rune(line) {
			var c rune
			switch char {
			case '.':
				c = '0'
			case '#':
				c = '1'
			default:
				fmt.Printf("ERROR ")
			}
			key := getKey(i, j)
			result[key] = c
		}
	}
	return result
}

func getImage(input []string) *aocImage {
	return &aocImage{
		algorithm:     getAlgorithm(input[0]),
		points:        getPoints(input[1]),
		infinityValue: '0',
		minI:          0,
		minJ:          0,
		maxI:          len(strings.Split(input[1], "\n")),
		maxJ:          len(strings.Split(input[1], "\n")[0]),
	}
}
