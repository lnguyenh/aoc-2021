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
	minX         int
	maxX         int
	minY         int
	maxY         int
	targetPoints map[string]bool

	// run trail
	points           map[string]bool
	maxHeightThisRun int
}

func getKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
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
	targetPoints := make(map[string]bool)
	for x := probe.minX; x <= probe.maxX; x++ {
		for y := probe.minY; y <= probe.maxY; y++ {
			targetPoints[getKey(x, y)] = true
		}
	}
	probe.targetPoints = targetPoints
	return &probe
}

func (probe *aocProbe) print(minX, maxX, minY, maxY int) {
	fmt.Printf("   ")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%v ", x%10)
	}
	fmt.Printf("\n")

	for y := maxY; y >= minY; y-- {
		fmt.Printf("%02v ", y%10)
		for x := minX; x <= maxX; x++ {
			key := getKey(x, y)
			if x == probe.x && y == probe.y {
				fmt.Printf("* ")
			} else if probe.points[key] {
				fmt.Printf("# ")
			} else if probe.targetPoints[key] {
				fmt.Printf("T ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (probe *aocProbe) doStep() {
	// position change
	probe.x = probe.x + probe.vX
	probe.y = probe.y + probe.vY

	// velocity change
	if probe.vX > 0 {
		probe.vX--
	} else if probe.vX < 0 {
		probe.vX++
	}
	probe.vY--

	// keep track of visited points and max height reached
	probe.points[getKey(probe.x, probe.y)] = true
	if probe.y > probe.maxHeightThisRun {
		probe.maxHeightThisRun = probe.y
	}

}

func (probe *aocProbe) isInTarget() bool {
	return probe.minX <= probe.x &&
		probe.x <= probe.maxX &&
		probe.minY <= probe.y &&
		probe.y <= probe.maxY
}

func (probe *aocProbe) isDead() bool {
	if probe.y <= probe.minY ||
		probe.vX < 0 && probe.x < probe.minX ||
		probe.vX > 0 && probe.x > probe.maxX ||
		probe.vX == 0 && !(probe.minX <= probe.x && probe.x <= probe.maxX) {
		return true
	}
	return false
}

func (probe *aocProbe) initializeRune(vX, vY int) {
	probe.x = 0
	probe.y = 0
	probe.vX = vX
	probe.vY = vY
	probe.points = make(map[string]bool)
	probe.points[getKey(0, 0)] = true
	probe.maxHeightThisRun = 0
}

func (probe *aocProbe) isVectorValid(vX, vY int) bool {
	probe.initializeRune(vX, vY)
	for {
		if probe.isInTarget() {
			return true
		}
		if probe.isDead() {
			break
		}
		probe.doStep()
	}
	return false
}

func (probe *aocProbe) scan() (int, int) {
	maxHeights := make([]int, 0, 2500)
	for vx := 0; vx < 300; vx++ {
		for vy := -300; vy < 300; vy++ {
			if probe.isVectorValid(vx, vy) {
				maxHeights = append(maxHeights, probe.maxHeightThisRun)
			}
		}
	}
	return utils.MaxSlice(maxHeights), len(maxHeights)
}

func Run(path string) {
	target := utils.StringSliceToIntSlice(utils.CleanSlice(utils.ReadFileAsStringSliceMulti(
		path, []string{"target area: ", ",", "x=", "y=", ".."})))
	probe := createProbe(target)
	maxHeight, numSuccesses := probe.scan()

	probe.isVectorValid(7, 2)
	probe.print(0, 30, -12, 5)

	fmt.Printf("Part 1 answer: %v\n", maxHeight)
	fmt.Printf("Part 2 answer: %v\n", numSuccesses)
}
