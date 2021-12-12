package day11

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type octopusGame struct {
	grid       [][]int
	workGrid   [][]int
	numFlashes int
	numSteps   int
	width      int
	length     int
}

func (game *octopusGame) resetWorkGrid() {
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			game.workGrid[i][j] = 0
		}
	}
}

func (game *octopusGame) isSynchronous() bool {
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			if game.grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func (game *octopusGame) copyWorkGridToMainGrid() {
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			game.grid[i][j] = game.workGrid[i][j]
		}
	}
}

func (game *octopusGame) applyStep() {
	game.resetWorkGrid()

	hasFlashed := make([][]bool, game.length)
	for i := range hasFlashed {
		hasFlashed[i] = make([]bool, game.width)
	}

	// plus one everything
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			game.workGrid[i][j] = game.grid[i][j] + 1
		}
	}

	// flash
	var numLoopFlashes int
	for {
		numLoopFlashes = 0
		for i := 0; i < game.length; i++ {
			for j := 0; j < game.width; j++ {
				if !hasFlashed[i][j] && game.workGrid[i][j] > 9 {
					// Flash !
					numLoopFlashes++
					game.numFlashes = game.numFlashes + 1

					// Spread
					spreads := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
						{i - 1, j - 1}, {i - 1, j + 1}, {i + 1, j - 1}, {i + 1, j + 1}}
					for _, spread := range spreads {
						y, x := spread[0], spread[1]
						if y >= 0 && y <= game.length-1 && x >= 0 && x <= game.width-1 {
							game.workGrid[y][x] = game.workGrid[y][x] + 1
						}
					}

					// Mark as flashed
					hasFlashed[i][j] = true
				}
			}
		}
		if numLoopFlashes == 0 {
			break
		}
	}

	// zero all the flashes
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			if game.workGrid[i][j] > 9 {
				game.workGrid[i][j] = 0
			}
		}
	}

	game.copyWorkGridToMainGrid()
	game.numSteps = game.numSteps + 1

}
func (game *octopusGame) print() {
	fmt.Printf("Main grid - Step #%v - Flashes %v\n", game.numSteps, game.numFlashes)
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			fmt.Printf("%v", game.grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (game *octopusGame) printWorkGrid() {
	fmt.Printf("Work grid - Step #%v - Flashes %v\n", game.numSteps, game.numFlashes)
	for i := 0; i < game.length; i++ {
		for j := 0; j < game.width; j++ {
			fmt.Printf("%v", game.workGrid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func createGame(input [][]int) octopusGame {
	grid := make([][]int, len(input))
	for i := range grid {
		grid[i] = make([]int, len(input[0]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			grid[i][j] = input[i][j]
		}
	}

	workGrid := make([][]int, len(input))
	for i := range workGrid {
		workGrid[i] = make([]int, len(input[0]))
	}

	return octopusGame{
		grid:       grid,
		workGrid:   workGrid,
		numFlashes: 0,
		numSteps:   0,
		width:      len(input[0]),
		length:     len(input),
	}
}

func doPart1(game octopusGame) int {
	for i := 0; i < 100; i++ {
		game.applyStep()
		// game.print()
	}
	return game.numFlashes
}

func doPart2(game octopusGame) int {
	for {
		if game.isSynchronous() {
			break
		}
		game.applyStep()
		// game.print()
	}
	return game.numSteps
}

func Run(path string) {
	input := utils.ReadFileAsSliceOfDigitIntSlices(path)
	game1 := createGame(input)
	game2 := createGame(input)
	answer1 := doPart1(game1)
	answer2 := doPart2(game2)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
