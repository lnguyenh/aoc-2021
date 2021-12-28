package day22

import (
	"github.com/lnguyenh/aoc-2021/utils"
	"sort"
)

type aocSlice struct {
	min int
	max int
}

type aocBlock struct {
	x aocSlice
	y aocSlice
	z aocSlice
}

func (b *aocBlock) getVolume() int {
	return (b.x.max - b.x.min) * (b.y.max - b.y.min) * (b.z.max - b.z.min)
}

type aocInstruction struct {
	isFull           bool
	cuboidBoundaries []int
}

type aocSpace struct {
	blocks map[aocBlock]bool
	xs     []int
	ys     []int
	zs     []int

	xMap map[int]int
	yMap map[int]int
	zMap map[int]int

	instructions []aocInstruction
}

func (s *aocSpace) getVolume() int {
	volume := 0
	for block, isFull := range s.blocks {
		if isFull {
			volume += block.getVolume()
		}
	}
	return volume
}

func (s *aocSpace) add(xMin, xMax, yMin, yMax, zMin, zMax int) {
	s.xs = append(s.xs, xMin, xMax+1)
	s.ys = append(s.ys, yMin, yMax+1)
	s.zs = append(s.zs, zMin, zMax+1)
}

func (s *aocSpace) buildAxes() {
	for _, instruction := range s.instructions {
		bounds := instruction.cuboidBoundaries
		s.add(bounds[0], bounds[1], bounds[2], bounds[3], bounds[4], bounds[5])
	}
}

func (s *aocSpace) applyInstructions() {
	for _, instruction := range s.instructions {
		bounds := instruction.cuboidBoundaries
		s.set(bounds[0], bounds[1], bounds[2], bounds[3], bounds[4], bounds[5], instruction.isFull)
	}
}

func (s *aocSpace) set(xMin, xMax, yMin, yMax, zMin, zMax int, isFull bool) {
	for i := s.xMap[xMin]; i < s.xMap[xMax+1]; i++ {
		for j := s.yMap[yMin]; j < s.yMap[yMax+1]; j++ {
			for k := s.zMap[zMin]; k < s.zMap[zMax+1]; k++ {
				x0, x1 := s.xs[i], s.xs[i+1]
				y0, y1 := s.ys[j], s.ys[j+1]
				z0, z1 := s.zs[k], s.zs[k+1]
				s.blocks[aocBlock{
					x: aocSlice{x0, x1},
					y: aocSlice{y0, y1},
					z: aocSlice{z0, z1},
				}] = isFull
			}
		}
	}
}

func (s *aocSpace) simplify() {
	s.xs = utils.IntSliceToSet(s.xs)
	s.ys = utils.IntSliceToSet(s.ys)
	s.zs = utils.IntSliceToSet(s.zs)
	sort.Ints(s.xs)
	sort.Ints(s.ys)
	sort.Ints(s.zs)

	for i, x := range s.xs {
		s.xMap[x] = i
	}
	for i, y := range s.ys {
		s.yMap[y] = i
	}
	for i, z := range s.zs {
		s.zMap[z] = i
	}

}

func (s *aocSpace) initializeGrid() {
	for i := 0; i < len(s.xs)-1; i++ {
		x0, x1 := s.xs[i], s.xs[i+1]
		for j := 0; j < len(s.ys)-1; j++ {
			y0, y1 := s.ys[j], s.ys[j+1]
			for k := 0; k < len(s.zs)-1; k++ {
				z0, z1 := s.zs[k], s.zs[k+1]
				s.blocks[aocBlock{
					x: aocSlice{x0, x1},
					y: aocSlice{y0, y1},
					z: aocSlice{z0, z1},
				}] = false
			}
		}
	}
}

func createSpace(instructions []aocInstruction) *aocSpace {
	space := aocSpace{
		blocks:       make(map[aocBlock]bool),
		xs:           make([]int, 0),
		ys:           make([]int, 0),
		zs:           make([]int, 0),
		xMap:         make(map[int]int),
		yMap:         make(map[int]int),
		zMap:         make(map[int]int),
		instructions: instructions,
	}
	return &space
}
