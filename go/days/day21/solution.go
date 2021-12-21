package day21

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

func doPart1(startingPositions []int) int {
	game := &aocGame1{
		scores:        []int{0, 0},
		positions:     []int{startingPositions[0], startingPositions[1]},
		currentPlayer: 0,
		dieValue:      1,
	}
	game.playCompleteGame()
	return game.getPart1Result()
}

func doPart2(startingPositions []int) int {
	game := &aocGame2{
		scoreCounts: make(map[gameRepresentation]int),
	}
	game.scoreCounts[gameRepresentation{
		score1:        0,
		score2:        0,
		position1:     startingPositions[0],
		position2:     startingPositions[1],
		currentPlayer: 0,
	}] = 1
	game.playCompleteGame()
	return utils.MaxSlice([]int{game.winner1Count, game.winner2Count})
}

func Run(path string) {
	input := utils.StringSliceToIntSlice(utils.CleanSlice(utils.ReadFileAsStringSliceMulti(path, []string{
		"\n",
		"Player 1 starting position: ",
		"Player 2 starting position: "})))
	fmt.Printf("input: %v\n", input)
	answer1 := doPart1(input)
	answer2 := doPart2(input)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
