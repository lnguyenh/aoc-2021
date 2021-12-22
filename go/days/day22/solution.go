package day22

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
	"regexp"
	"sort"
)

type aocSpace struct {
	cuboids []*aocCuboid
	cleans  map[string]*aocCuboid
}

func (s *aocSpace) getVolume() int {
	volume := 0
	for _, cuboid := range s.cuboids {
		volume += cuboid.getVolume()
	}
	return volume
}

func (s *aocSpace) save(newC *aocCuboid) {
	s.cuboids = append(s.cuboids, newC)
}

func (s *aocSpace) runChecks() {
	noIntersect := make([]int, 0)
IntersectLoop:
	for i, c := range s.cuboids {
		for _, c2 := range s.cuboids {
			if c == c2 {
				continue
			}
			fmt.Printf("toto\n")
			if c.intersects(c2) {
				continue IntersectLoop
			}
		}
		noIntersect = append(noIntersect, i)
	}
	for i, c := range s.cuboids {
		for j, c2 := range s.cuboids[:i] {
			if c2.contains(c) {
				fmt.Printf("%v contains %v\n", j, i)
			} else {
				// fmt.Printf("%v has %v points in %v   || %v\n", i, len(c.pointsIn(c2)), j, c.pointsIn(c2))
			}

		}
	}
}

func (s *aocSpace) makeClean() {
	for _, c := range s.cuboids {
		s.add(c)
	}
}

func (s *aocSpace) add(newC *aocCuboid) {
	for _, oldC := range s.cuboids {
		// no intersection
		if !oldC.intersects(newC) {
			continue
		}
		if oldC.contains(newC) {
			continue
		}

		newCPointsInOldC := newC.pointsIn(oldC)
		oldCPointsInNewC := oldC.pointsIn(newC)

		if len(newCPointsInOldC) >= len(oldCPointsInNewC) {
			// deal with newC in olC
			switch len(newCPointsInOldC) {
			case 1:
				if newC.b1.in(oldC) {
					s.add(createCuboid(newC.xMin, oldC.xMin-1, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, newC.yMax, oldC.zMax+1, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, oldC.yMax+1, newC.yMax, newC.zMin, oldC.zMax, true))
				} else if newC.b2.in(oldC) {
					s.add(createCuboid(oldC.xMax+1, newC.xMax, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, newC.yMax, oldC.zMax+1, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, oldC.yMax+1, newC.yMax, newC.zMin, oldC.zMax, true))
				} else if newC.b3.in(oldC) {
					s.add(createCuboid(oldC.xMax+1, newC.xMax, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, newC.yMax, oldC.zMax+1, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, oldC.yMin-1, newC.zMin, oldC.zMax, true))
				} else if newC.b4.in(oldC) {
					s.add(createCuboid(newC.xMin, oldC.xMin-1, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, newC.yMax, oldC.zMax+1, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, oldC.yMin-1, newC.zMin, oldC.zMax, true))
				} else if newC.t1.in(oldC) {
					s.add(createCuboid(newC.xMin, oldC.xMin-1, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, newC.yMax, newC.zMin, oldC.zMax-1, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, oldC.yMax+1, newC.yMax, oldC.zMax, newC.zMax, true))
				} else if newC.t2.in(oldC) {
					s.add(createCuboid(oldC.xMax+1, newC.xMax, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, newC.yMax, newC.zMin, oldC.zMax-1, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, oldC.yMax+1, newC.yMax, oldC.zMax, newC.zMax, true))
				} else if newC.t3.in(oldC) {
					s.add(createCuboid(oldC.xMax+1, newC.xMax, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, newC.yMax, newC.zMin, oldC.zMin-1, true))
					s.add(createCuboid(newC.xMin, oldC.xMax, newC.yMin, oldC.yMin-1, oldC.zMin, newC.zMax, true))
				} else if newC.t4.in(oldC) {
					s.add(createCuboid(newC.xMin, oldC.xMin-1, newC.yMin, newC.yMax, newC.zMin, newC.zMax, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, newC.yMax, newC.zMin, oldC.zMin-1, true))
					s.add(createCuboid(oldC.xMin, newC.xMax, newC.yMin, oldC.yMin-1, oldC.zMin, newC.zMax, true))
				}
			case 2:
				if newC.b1.in(oldC) && newC.b2.in(oldC) {

				} else if newC.b2.in(oldC) && newC.b3.in(oldC) {

				} else if newC.b3.in(oldC) && newC.b4.in(oldC) {

				} else if newC.b4.in(oldC) && newC.b1.in(oldC) {

				} else if newC.t1.in(oldC) && newC.t2.in(oldC) {

				} else if newC.t2.in(oldC) && newC.t3.in(oldC) {

				} else if newC.t3.in(oldC) && newC.t4.in(oldC) {

				} else if newC.t4.in(oldC) && newC.t1.in(oldC) {

				}
			case 4:
				if newC.b1.in(oldC) && newC.b2.in(oldC) && newC.t1.in(oldC) && newC.t2.in(oldC) {

				} else if newC.b2.in(oldC) && newC.t2.in(oldC) && newC.b3.in(oldC) && newC.t3.in(oldC) {

				} else if newC.b4.in(oldC) && newC.t4.in(oldC) && newC.b3.in(oldC) && newC.t3.in(oldC) {

				} else if newC.b4.in(oldC) && newC.t4.in(oldC) && newC.b1.in(oldC) && newC.t1.in(oldC) {

				} else if newC.b1.in(oldC) && newC.b2.in(oldC) && newC.b3.in(oldC) && newC.b4.in(oldC) {

				} else if newC.t1.in(oldC) && newC.t2.in(oldC) && newC.t3.in(oldC) && newC.t4.in(oldC) {

				}
			}
		} else {
			// deal with oldC in newC partially
		}

		// deal with oldC in newC totally

	}

	// no intersection to deal with
	s.cleans[newC.name()] = newC
}

type aocPoint struct {
	x int
	y int
	z int
}

func (p *aocPoint) in(c *aocCuboid) bool {
	return p.x >= c.xMin && p.x <= c.xMax &&
		p.y >= c.xMin && p.y <= c.yMax &&
		p.z >= c.zMin && p.z <= c.zMax
}

type aocCuboid struct {
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	zMin   int
	zMax   int
	isFull bool
	t1     aocPoint
	t2     aocPoint
	t3     aocPoint
	t4     aocPoint
	b1     aocPoint
	b2     aocPoint
	b3     aocPoint
	b4     aocPoint
}

func createCuboid(xMin, xMax, yMin, yMax, zMin, zMax int, isFull bool) *aocCuboid {
	c := aocCuboid{
		xMin, xMax,
		yMin, yMax,
		zMin, zMax,
		isFull,
		aocPoint{xMax, yMin, zMax},
		aocPoint{xMin, yMin, zMax},
		aocPoint{xMin, yMax, zMax},
		aocPoint{xMin, yMax, zMax},
		aocPoint{xMin, yMax, zMax},
		aocPoint{xMin, yMin, zMin},
		aocPoint{xMin, yMax, zMin},
		aocPoint{xMax, yMax, zMin}}
	return &c
}

func (c *aocCuboid) intersects(otherC *aocCuboid) bool {
	return c.xMin <= otherC.xMax && c.xMax >= otherC.xMin &&
		c.yMin <= otherC.yMax && c.yMax >= otherC.yMin &&
		c.zMin <= otherC.zMax && c.zMax >= otherC.zMin
}

func (c *aocCuboid) pointsIn(otherC *aocCuboid) map[aocPoint]bool {
	result := map[aocPoint]bool{}
	for _, point := range []aocPoint{c.t1, c.t2, c.t3, c.t4, c.b1, c.b2, c.b3, c.b4} {
		if point.in(otherC) {
			result[point] = true
		}
	}
	return result
}

func (c *aocCuboid) contains(otherC *aocCuboid) bool {
	return otherC.xMin >= c.xMin && otherC.xMax <= c.xMax &&
		otherC.yMin >= c.xMin && otherC.yMax <= c.yMax &&
		otherC.zMin >= c.zMin && otherC.zMax <= c.zMax
}
func (c *aocCuboid) getVolume() int {
	return (c.xMax - c.xMin) * (c.yMax - c.yMin) * (c.zMax - c.zMin)
}

func (c *aocCuboid) name() string {
	return fmt.Sprintf("x(%v.%v)y(%v.%v)z(%v.%v)", c.xMin, c.xMax, c.yMin, c.yMax, c.zMin, c.zMax)
}

func doPart2(instructions []string) int {
	space := aocSpace{
		cuboids: make([]*aocCuboid, 0),
		cleans:  make(map[string]*aocCuboid),
	}
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
		switch onOrOff {
		case "on":
			space.save(createCuboid(xs[0], xs[1], ys[0], ys[1], zs[0], zs[1], true))
		case "off":
			space.save(createCuboid(xs[0], xs[1], ys[0], ys[1], zs[0], zs[1], false))
		}
	}
	space.makeClean()

	return space.getVolume()
}

func Run(path string) {
	input := utils.ReadFileAsStringSlice(path, "\n")
	answer1 := doPart1(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	answer2 := doPart2(input)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
