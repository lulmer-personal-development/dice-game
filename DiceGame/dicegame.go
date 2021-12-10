package dicegame

import (
	"math/rand"
	"time"
)

func GameRoll(numOfDice int) int {
	if numOfDice == 0 {
		return 0
	}

	dice := make([]int, numOfDice)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOfDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}

	score, newDiceCount := calculateScoreAndNewDiceCount(dice, numOfDice)
	// score := 7
	// newDiceCount := numOfDice
	// for _, diceVal := range dice {
	// 	if diceVal == 3 {
	// 		score = 0
	// 		newDiceCount--
	// 	} else if score != 0 && diceVal < score {
	// 		score = diceVal
	// 	}
	// }
	// if score != 0 {
	// 	newDiceCount--
	// }

	return score + GameRoll(newDiceCount)
}

func calculateScoreAndNewDiceCount(dice []int, numOfDice int) (int, int) {
	score := 7
	newDiceCount := numOfDice
	for _, diceVal := range dice {
		if diceVal == 3 {
			score = 0
			newDiceCount--
		} else if score != 0 && diceVal < score {
			score = diceVal
		}
	}
	if score != 0 {
		newDiceCount--
	}

	return score, newDiceCount
}
