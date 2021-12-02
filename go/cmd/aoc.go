package main

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/days"
	"os"
)

func main() {

	args := os.Args[1:]
	day := args[0]
	input := args[1]
	inputPath := fmt.Sprintf("../inputs/%s/%s.txt", day, input)
	fmt.Printf("Running AOC for day %s with input '%s'\n", day, inputPath)
	days.Run(day, inputPath)
}
