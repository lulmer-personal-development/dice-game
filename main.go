package main

import (
	"fmt"
	"sort"
	"time"

	"FCC-challenge/dicegame"
)

func sortScoreMapKeys(scoreMap map[int]int) []int {
	scores := []int{}
	for sc := range scoreMap {
		scores = append(scores, sc)
	}

	sort.Ints(scores)

	return scores
}

func main() {
	const gameIterations int = 10000
	scoreMap := make(map[int]int)

	startTime := time.Now()
	for i := 0; i < gameIterations; i++ {
		newScore := dicegame.GameRoll(5)
		scoreMap[newScore]++
	}
	elapsedTime := time.Since(startTime)

	scores := sortScoreMapKeys(scoreMap)

	for _, score := range scores {
		percentage := float64(scoreMap[score]) / float64(gameIterations)
		fmt.Printf("Score: %d occurs %.4f occurred %d times\n", score, percentage, scoreMap[score])
	}
	fmt.Printf("Total simulation time took %f seconds\n", elapsedTime.Seconds())

}
