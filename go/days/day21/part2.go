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
	rollCount     int
	rollTmpValue  int
}

func getNewPosition(currentPosition, rollValue int) int {
	newPosition := currentPosition + rollValue
	if newPosition > 10 {
		return newPosition % 10
	}
	return newPosition
}

func (r *gameRepresentation) playOnce() [3]gameRepresentation {
	rollCount := r.rollCount + 1
	if rollCount > 2 {
		// move pawns and reset rollCount
		if r.currentPlayer == 0 {
			positionA := getNewPosition(r.position1, r.rollTmpValue+1)
			positionB := getNewPosition(r.position1, r.rollTmpValue+2)
			positionC := getNewPosition(r.position1, r.rollTmpValue+3)
			return [3]gameRepresentation{
				{
					score1:        r.score1 + positionA,
					score2:        r.score2,
					position1:     positionA,
					position2:     r.position2,
					currentPlayer: 1,
					rollCount:     0,
					rollTmpValue:  0,
				},
				{
					score1:        r.score1 + positionB,
					score2:        r.score2,
					position1:     positionB,
					position2:     r.position2,
					currentPlayer: 1,
					rollCount:     0,
					rollTmpValue:  0,
				},
				{
					score1:        r.score1 + positionC,
					score2:        r.score2,
					position1:     positionC,
					position2:     r.position2,
					currentPlayer: 1,
					rollCount:     0,
					rollTmpValue:  0,
				},
			}
		} else {
			positionA := getNewPosition(r.position2, r.rollTmpValue+1)
			positionB := getNewPosition(r.position2, r.rollTmpValue+2)
			positionC := getNewPosition(r.position2, r.rollTmpValue+3)
			return [3]gameRepresentation{
				{
					score1:        r.score1,
					score2:        r.score2 + positionA,
					position1:     r.position1,
					position2:     positionA,
					currentPlayer: 0,
					rollCount:     0,
					rollTmpValue:  0,
				},
				{
					score1:        r.score1,
					score2:        r.score2 + positionB,
					position1:     r.position1,
					position2:     positionB,
					currentPlayer: 0,
					rollCount:     0,
					rollTmpValue:  0,
				},
				{
					score1:        r.score1,
					score2:        r.score2 + positionC,
					position1:     r.position1,
					position2:     positionC,
					currentPlayer: 0,
					rollCount:     0,
					rollTmpValue:  0,
				},
			}
		}
	} else {
		// Keep rolling and just update the temporary roll accumulator and count
		if r.currentPlayer == 0 {
			return [3]gameRepresentation{
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 1,
				},
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 2,
				},
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 3,
				},
			}
		} else {
			return [3]gameRepresentation{
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 1,
				},
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 2,
				},
				{
					score1:        r.score1,
					score2:        r.score2,
					position1:     r.position1,
					position2:     r.position2,
					currentPlayer: r.currentPlayer,
					rollCount:     rollCount,
					rollTmpValue:  r.rollTmpValue + 3,
				},
			}
		}
	}
}

func (game *aocGame2) step() {
	newScoreCounts := make(map[gameRepresentation]int)
	for representation, count := range game.scoreCounts {
		newRs := representation.playOnce()
		for _, r := range newRs {
			if r.score1 >= 21 {
				game.winner1Count += count
			} else if r.score2 >= 21 {
				game.winner2Count += count
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
