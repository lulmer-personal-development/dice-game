package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func gameRoll(numOfDice int) int {
	if numOfDice == 0 {
		return 0
	}

	dice := make([]int, numOfDice)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOfDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}

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

	return score + gameRoll(newDiceCount)
}

func sortScoreMapKeys(scoreMap map[int]int) []int {
	scores := []int{}
	for sc := range scoreMap {
		scores = append(scores, sc)
	}

	sort.Ints(scores)

	return scores
}

func main() {
	scoreMap := make(map[int]int)

	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		newScore := gameRoll(5)
		scoreMap[newScore]++
	}
	elapsedTime := time.Since(startTime)

	scores := sortScoreMapKeys(scoreMap)

	for _, score := range scores {
		percentage := float64(scoreMap[score]) / 100.0
		fmt.Println(fmt.Sprintf("Score: %d occurs %.2f occurred %d times", score, percentage, scoreMap[score]))
	}
	fmt.Print(fmt.Sprintf("Total simulation time took %f seconds", elapsedTime.Seconds()))

}
