package utils

import (
	"fmt"
	"strconv"
)

func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func DigitToRune(d int) rune {
	return rune(d) + '0'
}

func StringToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func BitsToInt(s string) int {
	if value, err := strconv.ParseInt(s, 2, 64); err != nil {
		fmt.Println(err)
		return 0
	} else {
		return int(value)
	}
}
