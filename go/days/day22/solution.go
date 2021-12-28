package day22

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
	"sort"
)

func doPart2(instructions []string) int {
	r, _ := regexp.Compile("^(\\w+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)$")
	var space *aocSpace
	space = createSpace()

	for _, instruction := range instructions {
		groups := r.FindStringSubmatch(instruction)
		_, x0, x1, y0, y1, z0, z1 := groups[1], groups[2], groups[3], groups[4], groups[5], groups[6], groups[7]
		xs := utils.StringSliceToIntSlice([]string{x0, x1})
		ys := utils.StringSliceToIntSlice([]string{y0, y1})
		zs := utils.StringSliceToIntSlice([]string{z0, z1})
		sort.Ints(xs)
		sort.Ints(ys)
		sort.Ints(zs)
		space.add(xs[0], xs[1], ys[0], ys[1], zs[0], zs[1])
	}
	space.simplify()
	// space.initializeGrid()
	for _, instruction := range instructions {
		groups := r.FindStringSubmatch(instruction)
		onOrOff, x0, x1, y0, y1, z0, z1 := groups[1], groups[2], groups[3], groups[4], groups[5], groups[6], groups[7]
		xs := utils.StringSliceToIntSlice([]string{x0, x1})
		ys := utils.StringSliceToIntSlice([]string{y0, y1})
		zs := utils.StringSliceToIntSlice([]string{z0, z1})
		sort.Ints(xs)
		sort.Ints(ys)
		sort.Ints(zs)
		space.set(xs[0], xs[1], ys[0], ys[1], zs[0], zs[1], onOrOff == "on")
	}

	return space.getVolume()
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	answer1 := doPart1(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := doPart2(input)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
