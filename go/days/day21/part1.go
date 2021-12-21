package day21

import "github.com/lnguyenh/aoc-2021/utils"

type aocGame1 struct {
	scores        []int
	positions     []int
	currentPlayer int
	dieValue      int
	numRolls      int
}

func (game *aocGame1) getRoll() int {
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

func (game *aocGame1) getNextPosition(currentPosition, rollValue int) int {
	newPosition := currentPosition + rollValue%10
	if newPosition > 10 {
		return newPosition % 10
	}
	return newPosition
}

func (game *aocGame1) playOnce() {
	rollValue := game.getRoll()
	currentPlayer := game.currentPlayer
	game.positions[currentPlayer] = game.getNextPosition(game.positions[currentPlayer], rollValue)
	game.scores[currentPlayer] = game.scores[currentPlayer] + game.positions[currentPlayer]
	game.currentPlayer = (currentPlayer + 1) % 2
	game.numRolls += 3
}

func (game *aocGame1) playCompleteGame() {
	for {
		if game.scores[0] >= 1000 || game.scores[1] >= 1000 {
			break
		}
		game.playOnce()
	}
}

func (game *aocGame1) getPart1Result() int {
	return utils.MinSlice(game.scores) * game.numRolls
}
