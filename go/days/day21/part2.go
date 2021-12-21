package day21

type aocGame2 struct {
	scoreCounts  map[gameRepresentation]int
	winner1Count int
	winner2Count int
}

type gameRepresentation struct {
	score1        int
	score2        int
	position1     int
	position2     int
	currentPlayer int
}

func getNewPosition(currentPosition, rollValue int) int {
	newPosition := currentPosition + rollValue
	if newPosition > 10 {
		return newPosition % 10
	}
	return newPosition
}

func (r *gameRepresentation) playOnce() [3]gameRepresentation {
	if r.currentPlayer == 0 {
		positionA := getNewPosition(r.position1, 1)
		positionB := getNewPosition(r.position1, 2)
		positionC := getNewPosition(r.position1, 3)
		return [3]gameRepresentation{
			{
				score1:        r.score1 + positionA,
				score2:        r.score2,
				position1:     positionA,
				position2:     r.position2,
				currentPlayer: 1,
			},
			{
				score1:        r.score1 + positionB,
				score2:        r.score2,
				position1:     positionB,
				position2:     r.position2,
				currentPlayer: 1,
			},
			{
				score1:        r.score1 + positionC,
				score2:        r.score2,
				position1:     positionC,
				position2:     r.position2,
				currentPlayer: 1,
			},
		}
	} else {
		positionA := getNewPosition(r.position2, 1)
		positionB := getNewPosition(r.position2, 2)
		positionC := getNewPosition(r.position2, 3)
		return [3]gameRepresentation{
			{
				score1:        r.score1,
				score2:        r.score2 + positionA,
				position1:     r.position1,
				position2:     positionA,
				currentPlayer: 0,
			},
			{
				score1:        r.score1,
				score2:        r.score2 + positionB,
				position1:     r.position1,
				position2:     positionB,
				currentPlayer: 0,
			},
			{
				score1:        r.score1,
				score2:        r.score2 + positionC,
				position1:     r.position1,
				position2:     positionC,
				currentPlayer: 0,
			},
		}
	}
}

func (game *aocGame2) step() {
	newScoreCounts := make(map[gameRepresentation]int)
	for representation, count := range game.scoreCounts {
		newRs := representation.playOnce()
		for _, r := range newRs {
			if r.score1 >= 21 {
				game.winner1Count += 1
			} else if r.score2 >= 21 {
				game.winner2Count += 1
			} else {
				newScoreCounts[r] += count
			}
		}
	}
	game.scoreCounts = newScoreCounts
}

func (game *aocGame2) playCompleteGame() {
	for {
		if len(game.scoreCounts) == 0 {
			break
		}
		game.step()
	}
}
