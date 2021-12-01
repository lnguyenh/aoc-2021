package main

import (
	"fmt"
	"go/days/01/go/days/day01"
	"os"
)

func main() {

	args := os.Args[1:]
	day := args[0]
	input := args[1]

	inputPath := fmt.Sprintf("inputs/%s/%s", day, input)

	fmt.Printf("Running AOC for day %s with input %s\n", day, inputPath)

	switch day {
	case "01":
		day01.Run(inputPath)
	default:
		fmt.Printf("Cannot find a solution for day %s", day)
	}

}
