package day17

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type aocProbe struct {
	// position
	x int
	y int

	// velocity
	vX int
	vY int

	// target
	minX int
	maxX int
	minY int
	maxY int
}

func createProbe(target []int) *aocProbe {
	probe := aocProbe{
		x:    0,
		y:    0,
		minX: utils.MinSlice(target[:2]),
		maxX: utils.MaxSlice(target[:2]),
		minY: utils.MinSlice(target[2:]),
		maxY: utils.MaxSlice(target[2:]),
	}
	return &probe
}

func (probe *aocProbe) doStep() {
	probe.x = probe.x + probe.vX
	probe.y = probe.y + probe.vY
	if probe.vX > 0 {
		probe.vX--
	} else if probe.vX < 0 {
		probe.vX++
	}
	probe.vY--
}

func doPart1(target []int) int {
	probe := createProbe(target)
	probe.doStep()
	return 0
}

func doPart2() int {
	return 0
}

func Run(path string) {
	target := utils.StringSliceToIntSlice(utils.CleanSlice(utils.ReadFileAsStringSliceMulti(
		path, []string{"target area: ", ",", "x=", "y=", ".."})))
	fmt.Printf("input: %v\n", target)
	answer1 := doPart1(target)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
