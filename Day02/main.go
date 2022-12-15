package main

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

const (
	ROCK = iota + 1
	PAPER
	SCISSORS
)

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0
)

//Player inside box is winning/losing
var WINNING_MOVE = map[int]int{ROCK: SCISSORS, SCISSORS: PAPER, PAPER: ROCK}
var LOSING_MOVE = map[int]int{SCISSORS: ROCK, PAPER: SCISSORS, ROCK: PAPER}

func main() {
	TotalScore := 0
	utils.ReadFile("Day02/Day2.txt", func(text string) {
		words := strings.Split(text, " ")
		OppChoice := ConvertOppChoice(words[0])
		MyChoice := ConvertMyChoice(words[1])
		RoundScore := MyChoice
		switch MyChoice {
		case OppChoice:
			RoundScore += DRAW
		case WINNING_MOVE[OppChoice]:
			RoundScore += LOSE
		default:
			RoundScore += WIN
		}
		TotalScore += RoundScore
	})
	fmt.Println("First Total Score:", TotalScore)

	TotalScore = 0
	utils.ReadFile("Day02/Day2.txt", func(text string) {
		RoundScore := 0
		words := strings.Split(text, " ")
		OppChoice := ConvertOppChoice(words[0])
		RoundEnd := ConvertRoundEnd(words[1])

		switch RoundEnd {
		case WIN:
			RoundScore = WIN + LOSING_MOVE[OppChoice]
		case LOSE:
			RoundScore = LOSE + WINNING_MOVE[OppChoice]
		default:
			RoundScore = DRAW + OppChoice
		}
		TotalScore += RoundScore
	})
	fmt.Println("Total Score:", TotalScore)

}

func ConvertRoundEnd(MyChoice string) int {
	switch MyChoice {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	case "Z":
		return WIN
	}
	return -1
}

func ConvertMyChoice(MyChoice string) int {
	switch MyChoice {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISSORS
	}
	return -1
}

func ConvertOppChoice(MyChoice string) int {
	switch MyChoice {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	}
	return -1
}
