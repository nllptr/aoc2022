package day2

import (
	"log"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) (int, int) {
	input := util.ReadFile(filename)
	parsed1 := Parse(input, asChoicePair)
	totalScore := GetTotalScore(parsed1)
	parsed2 := Parse(input, asCombination)
	totalScore2 := GetTotalScore(parsed2)
	return totalScore, totalScore2
}

type scorable interface {
	score() int
}

type choicePair struct {
	opponent choice
	you      choice
}

func (c choicePair) score() int {
	outcome := getOutcome(c.opponent, c.you)
	return int(c.you) + int(outcome)
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

func (c combination) score() int {
	return int(getMatch(c)) + int(c.outcome)
}

type outcome int

const (
	LOSS outcome = 0
	DRAW outcome = 3
	WIN  outcome = 6
)

func Parse(input []byte, parseFunc func([]string) scorable) []scorable {
	asString := string(input)
	rows := strings.Split(asString, "\n")
	parsed := []scorable{}
	for _, row := range rows {
		columns := strings.Split(row, " ")
		combination := parseFunc(columns)
		parsed = append(parsed, combination)
	}
	return parsed
}

func asChoicePair(input []string) scorable {
	return choicePair{asChoice(input[0]), asChoice(input[1])}
}

func asCombination(input []string) scorable {
	return combination{asChoice(input[0]), asOutcome(input[1])}
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

func getOutcome(opponent, you choice) outcome {
	if (opponent == ROCK && you == PAPER) || (opponent == PAPER && you == SCISSORS) || (opponent == SCISSORS && you == ROCK) {
		return WIN
	}
	if (opponent == ROCK && you == SCISSORS) || (opponent == PAPER && you == ROCK) || (opponent == SCISSORS && you == PAPER) {
		return LOSS
	}
	return DRAW
}

func GetTotalScore(combinations []scorable) int {
	score := 0
	for _, combination := range combinations {
		score = score + combination.score()
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
