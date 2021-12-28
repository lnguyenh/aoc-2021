package utils

import (
	"sort"
	"strings"
)

//
// Int
//

func SumSlice(measurement []int) int {
	sum := 0
	for _, value := range measurement {
		sum += value
	}
	return sum
}

func MaxSlice(slice []int) int {
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func MinSlice(slice []int) int {
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntSliceToSet(list []int) []int {
	var set []int
	for _, value := range list {
		if !IntInSlice(value, set) {
			set = append(set, value)
		}
	}
	return set
}

func InsertInIntSlice(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

//
// Runes
//

func RuneInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SortRuneSlice(runeSlice []rune) []rune {
	r := runeSlice[:]
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}

func RuneRevertSlice(list []rune) []rune {
	maxIndex := len(list) - 1
	reverted := make([]rune, len(list))
	for i, value := range list {
		reverted[maxIndex-i] = value
	}
	return reverted
}

func RuneSliceToSet(list []rune) []rune {
	var set []rune
	for _, value := range list {
		if !RuneInSlice(value, set) {
			set = append(set, value)
		}
	}
	return set
}

//
// Strings
//

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringsNotInSlice(candidates []string, list []string) []string {
	results := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		if !StringInSlice(candidate, list) {
			results = append(results, candidate)
		}
	}
	return results
}

func CleanSlice(slice []string) []string {
	result := make([]string, 0, len(slice))
	for _, word := range slice {
		trimmed := strings.TrimSpace(word)
		if len(trimmed) > 0 {
			result = append(result, trimmed)
		}
	}
	return result
}

func StringSliceToIntSlice(list []string) []int {
	result := make([]int, 0, len(list))
	for _, element := range list {
		result = append(result, StringToInt(element))
	}
	return result
}
