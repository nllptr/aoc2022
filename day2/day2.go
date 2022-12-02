package day2

import (
	"log"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) (int, int) {
	input := util.ReadFile(filename)
	parsed1 := parse1(input)
	totalScore := getTotalScore1(parsed1)
	parsed2 := parse2(input)
	totalScore2 := getTotalScore2(parsed2)
	return totalScore, totalScore2
}

type choicePair struct {
	opponent choice
	you      choice
}

type choice int

const (
	ROCK     choice = 1
	PAPER    choice = 2
	SCISSORS choice = 3
)

type combination struct {
	choice  choice
	outcome outcome
}

type outcome int

const (
	LOSS outcome = 0
	DRAW outcome = 3
	WIN  outcome = 6
)

func parse1(input []byte) []choicePair {
	asString := string(input)
	rows := strings.Split(asString, "\n")
	parsed := []choicePair{}
	for _, row := range rows {
		columns := strings.Split(row, " ")
		combination := choicePair{asChoice(columns[0]), asChoice(columns[1])}
		parsed = append(parsed, combination)
	}
	return parsed
}

func parse2(input []byte) []combination {
	asString := string(input)
	rows := strings.Split(asString, "\n")
	parsed := []combination{}
	for _, row := range rows {
		columns := strings.Split(row, " ")
		combination := combination{asChoice(columns[0]), asOutcome(columns[1])}
		parsed = append(parsed, combination)
	}
	return parsed
}

func asChoice(input string) choice {
	if !strings.Contains("ABCXYZ", input) {
		log.Fatal("invalid choice in input:", input)
	}
	var asChoice choice
	switch input {
	case "A", "X":
		asChoice = ROCK
	case "B", "Y":
		asChoice = PAPER
	case "C", "Z":
		asChoice = SCISSORS
	default:
		log.Fatal("no matching choice for", input)
	}
	return asChoice
}

func asOutcome(input string) outcome {
	if !strings.Contains("XYZ", input) {
		log.Fatal("invalid input:", input)
	}
	var asOutcome outcome
	switch input {
	case "X":
		asOutcome = LOSS
	case "Y":
		asOutcome = DRAW
	case "Z":
		asOutcome = WIN
	}
	return asOutcome
}

func getScore(c choicePair) int {
	outcome := getOutcome(c.opponent, c.you)
	return int(c.you) + int(outcome)
}

func getOutcome(opponent, you choice) outcome {
	if (opponent == ROCK && you == PAPER) || (opponent == PAPER && you == SCISSORS) || (opponent == SCISSORS && you == ROCK) {
		return WIN
	}
	if (opponent == ROCK && you == SCISSORS) || (opponent == PAPER && you == ROCK) || (opponent == SCISSORS && you == PAPER) {
		return LOSS
	}
	return DRAW
}

func getTotalScore1(input []choicePair) int {
	score := 0
	for _, combination := range input {
		score = score + getScore(combination)
	}
	return score
}

func getMatch(combination combination) choice {
	switch combination.outcome {
	case WIN:
		return getWinnerFor(combination.choice)
	case LOSS:
		return getLoserFor(combination.choice)
	default:
		return combination.choice
	}
}

func getWinnerFor(c choice) choice {
	var winner choice
	switch c {
	case ROCK:
		winner = PAPER
	case PAPER:
		winner = SCISSORS
	case SCISSORS:
		winner = ROCK
	}
	return winner
}

func getLoserFor(c choice) choice {
	var loser choice
	switch c {
	case ROCK:
		loser = SCISSORS
	case PAPER:
		loser = ROCK
	case SCISSORS:
		loser = PAPER
	}
	return loser
}

func getTotalScore2(combinations []combination) int {
	score := 0
	for _, combination := range combinations {
		score = score + getCombinationScore(combination)
	}
	return score
}

func getCombinationScore(c combination) int {
	return int(getMatch(c)) + int(choice(c.outcome))
}
