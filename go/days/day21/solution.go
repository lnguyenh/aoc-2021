package day21

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

type aocGame struct {
	scores        []int
	positions     []int
	currentPlayer int
	dieValue      int
	numRolls      int
}

func (game *aocGame) getRoll() int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += game.dieValue
		game.dieValue = game.dieValue + 1
		if game.dieValue > 100 {
			game.dieValue = 1
		}
	}
	return sum
}

func (game *aocGame) getNextPosition(currentPosition, rollValue int) int {
	newPosition := currentPosition + rollValue%10
	if newPosition > 10 {
		return newPosition % 10
	}
	return newPosition
}

func (game *aocGame) playOnce() {
	rollValue := game.getRoll()
	currentPlayer := game.currentPlayer
	game.positions[currentPlayer] = game.getNextPosition(game.positions[currentPlayer], rollValue)
	game.scores[currentPlayer] = game.scores[currentPlayer] + game.positions[currentPlayer]
	game.currentPlayer = (currentPlayer + 1) % 2
	game.numRolls += 3
}

func (game *aocGame) playCompleteGame() {
	for {
		if game.scores[0] >= 1000 || game.scores[1] >= 1000 {
			break
		}
		game.playOnce()
	}
}

func (game *aocGame) getPart1Result() int {
	return utils.MinSlice(game.scores) * game.numRolls
}

func doPart1(startingPositions []int) int {
	game := &aocGame{
		scores:        []int{0, 0},
		positions:     []int{startingPositions[0], startingPositions[1]},
		currentPlayer: 0,
		dieValue:      1,
	}

	// x := game.getNextPosition(10, 5)
	// fmt.Println(x)

	game.playCompleteGame()
	return game.getPart1Result()
}

func doPart2() int {
	return 0
}

func Run(path string) {
	input := utils.StringSliceToIntSlice(utils.CleanSlice(utils.ReadFileAsStringSliceMulti(path, []string{
		"\n",
		"Player 1 starting position: ",
		"Player 2 starting position: "})))
	fmt.Printf("input: %v\n", input)
	answer1 := doPart1(input)
	answer2 := doPart2()
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}
