package main

import (
	"fmt"


	utils "github.com/rustyb/aoc-go-2021/2022"
)


func main() {
	rounds := utils.ReadFile("\n", "./day-02-input.txt")

	scores := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	scoresPartB := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}
	totalScore := 0
	totalScoreB := 0
	for _, ele := range rounds { // you can escape index by _ keyword
		totalScore += scores[ele]
		totalScoreB += scoresPartB[ele]
	}
	fmt.Printf("Total score %d \n",  totalScore)
	fmt.Printf("Total score part 2 %d \n",  totalScoreB)
}
