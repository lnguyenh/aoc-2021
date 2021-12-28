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
	return (b.x.max + 1 - b.x.min) * (b.y.max + 1 - b.y.min) * (b.z.max + 1 - b.z.min)
}

type aocSpace struct {
	blocks map[aocBlock]bool
	xs     []int
	ys     []int
	zs     []int
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
	s.xs = append(s.xs, xMin, xMax)
	s.ys = append(s.ys, yMin, yMax)
	s.zs = append(s.zs, zMin, zMax)
}

func (s *aocSpace) set(xMin, xMax, yMin, yMax, zMin, zMax int, isFull bool) {
	for i := 0; i < len(s.xs)-1; i++ {
		x0, x1 := s.xs[i], s.xs[i+1]
		for j := 0; j < len(s.ys)-1; j++ {
			y0, y1 := s.ys[j], s.ys[j+1]
			for k := 0; k < len(s.zs)-1; k++ {
				z0, z1 := s.zs[k], s.zs[k+1]

				if xMin <= x0 && xMax >= x1 && yMin <= y0 && yMax >= y1 && zMin <= z0 && zMax >= z1 {
					s.blocks[aocBlock{
						x: aocSlice{x0, x1},
						y: aocSlice{y0, y1},
						z: aocSlice{z0, z1},
					}] = isFull
				}

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

func createSpace() *aocSpace {
	space := aocSpace{
		blocks: make(map[aocBlock]bool),
		xs:     make([]int, 0),
		ys:     make([]int, 0),
		zs:     make([]int, 0),
	}
	return &space
}
