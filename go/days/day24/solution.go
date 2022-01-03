package day24

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"strings"
)

type aocAlu struct {
	w          int
	x          int
	y          int
	z          int
	input      []int
	inputIndex int
	program    []string
}

func (alu *aocAlu) reset(input []int) {
	alu.w = 0
	alu.x = 0
	alu.y = 0
	alu.z = 0
	alu.inputIndex = 0
	alu.input = input
}

// isValid, isError
func (alu *aocAlu) isValid(input []int) (bool, bool, bool) {
	if utils.IntInSlice(0, input) {
		return false, false, true
	}
	alu.reset(input)
	for _, line := range alu.program {
		if !alu.doInstruction(strings.Split(line, " ")) {
			return false, true, false
		}
	}
	return alu.z == 0, false, false
}

func (alu *aocAlu) isValidInt(input int) (bool, bool, bool) {
	asIntSlice := make([]int, 0, 14)
	asString := fmt.Sprintf("%v", input)
	asRuneSlice := []rune(asString)
	for _, r := range asRuneSlice {
		asIntSlice = append(asIntSlice, utils.RuneToInt(r))
	}
	return alu.isValid(asIntSlice)
}

func (alu *aocAlu) getValue(parameter string) int {
	switch parameter {
	case "w":
		return alu.w
	case "x":
		return alu.x
	case "y":
		return alu.y
	case "z":
		return alu.z
	default:
		return utils.StringToInt(parameter)
	}
}

func (alu *aocAlu) save(variable string, value int) {
	switch variable {
	case "w":
		alu.w = value
	case "x":
		alu.x = value
	case "y":
		alu.y = value
	case "z":
		alu.z = value
	default:
		fmt.Printf("ERROR unknown variable\n")
	}
}

func (alu *aocAlu) getNextInput() int {
	inputIndex := alu.inputIndex
	alu.inputIndex += 1
	return alu.input[inputIndex]
}

func (alu *aocAlu) doInstruction(instruction []string) bool {
	defer func() bool {
		if err := recover(); err != nil {
			return false
		}
		return true
	}()

	keyword, variable := instruction[0], instruction[1]
	a := alu.getValue(variable)
	switch keyword {
	case "inp":
		alu.save(variable, alu.getNextInput())
	case "add":
		b := alu.getValue(instruction[2])
		alu.save(variable, a+b)
	case "mul":
		b := alu.getValue(instruction[2])
		alu.save(variable, a*b)
	case "div":
		b := alu.getValue(instruction[2])
		if b == 0 {
			// fmt.Printf("division by zero\n")
			panic("division by zero\n")
		}
		alu.save(variable, a/b)
	case "mod":
		b := alu.getValue(instruction[2])
		if a < 0 || b <= 0 {
			// fmt.Printf("modulo error")
			panic("modulo error")
		}
		alu.save(variable, a%b)
	case "eql":
		b := alu.getValue(instruction[2])
		if a != b {
			alu.save(variable, 0)
		} else {
			alu.save(variable, 1)
		}
	default:
		fmt.Printf("ERROR unknown instruction\n")
	}
	return true
}

func doPart1(program []string) int {
	// base := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	alu := aocAlu{program: program}
	number := 13579246899999
	for {
		isValid, isError, hasZeroes := alu.isValidInt(number)
		if !isError {
			//fmt.Printf("error for %v", number)
		}
		if isValid || number < 0 {
			break
		}
		if !hasZeroes && alu.z < 1000 {
			fmt.Printf("%v: z is %v\n", number, alu.z)
		}
		number -= 1
	}

	// alu.isValid([]int{4})
	return number
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	answer1 := doPart1(input)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
