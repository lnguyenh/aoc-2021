package dayXX

import (
	"fmt"
	"go_aoc/filereader"
)

func Part1(path string) {
	values := filereader.ReadAsIntArray(path)
	fmt.Println(values)
}

func Part2(path string) {
	values := filereader.ReadAsIntArray(path)
	fmt.Println(values)
}

func Run(path string) {
	Part1(path)
	Part2(path)
}