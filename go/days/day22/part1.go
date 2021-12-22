package day22

import (
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
	"sort"
)

type aocCube struct {
	x int
	y int
	z int
}

func doPart1(instructions []string) int {
	cubeStatuses := make(map[aocCube]bool)
	r, _ := regexp.Compile("^(\\w+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)$")
	for _, instruction := range instructions {
		groups := r.FindStringSubmatch(instruction)
		onOrOff, x0, x1, y0, y1, z0, z1 := groups[1], groups[2], groups[3], groups[4], groups[5], groups[6], groups[7]
		xs := utils.StringSliceToIntSlice([]string{x0, x1})
		ys := utils.StringSliceToIntSlice([]string{y0, y1})
		zs := utils.StringSliceToIntSlice([]string{z0, z1})
		sort.Ints(xs)
		sort.Ints(ys)
		sort.Ints(zs)
		xs[0] = utils.MaxSlice([]int{xs[0], -50})
		ys[0] = utils.MaxSlice([]int{ys[0], -50})
		zs[0] = utils.MaxSlice([]int{zs[0], -50})
		xs[1] = utils.MinSlice([]int{xs[1], 50})
		ys[1] = utils.MinSlice([]int{ys[1], 50})
		zs[1] = utils.MinSlice([]int{zs[1], 50})

		statusToSet := onOrOff == "on"
		for x := xs[0]; x <= xs[1]; x++ {
			for y := ys[0]; y <= ys[1]; y++ {
				for z := zs[0]; z <= zs[1]; z++ {
					cubeStatuses[aocCube{x: x, y: y, z: z}] = statusToSet
				}
			}
		}
	}
	count := 0
	for _, status := range cubeStatuses {
		if status {
			count++
		}
	}
	return count
}
