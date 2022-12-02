package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAsChoicepairInput(t *testing.T) {
	input := []byte("A Y\nB X\nC Z")
	expected := []scorable{
		choicePair{ROCK, PAPER},
		choicePair{PAPER, ROCK},
		choicePair{SCISSORS, SCISSORS},
	}

	parsed := parse(input, asChoicePair)
	assert.Equal(t, expected, parsed, "parsed does not contain all expected values")
}
func TestParseAsCombinationInput(t *testing.T) {
	input := []byte("A Y\nB X\nC Z")
	expected := []scorable{
		combination{ROCK, DRAW},
		combination{PAPER, LOSS},
		combination{SCISSORS, WIN},
	}
	parsed := parse(input, asCombination)
	assert.Equal(t, expected, parsed, "parsed does not contain all expected values")
}

func TestGetScore(t *testing.T) {
	cases := []struct {
		input    choicePair
		expected int
	}{
		{choicePair{ROCK, PAPER}, 8},
		{choicePair{PAPER, ROCK}, 1},
		{choicePair{SCISSORS, SCISSORS}, 6},
	}
	for _, c := range cases {
		score := c.input.score()
		if score != c.expected {
			t.Error("test failed. expected", c.expected, "but got", score)
			t.Fail()
		}
	}
}

func TestGetTotalScore1(t *testing.T) {
	input := []scorable{
		choicePair{ROCK, PAPER},
		choicePair{PAPER, ROCK},
		choicePair{SCISSORS, SCISSORS},
	}
	expected := 15
	score := getTotalScore(input)
	if score != expected {
		t.Error("test failed. expected", expected, "but got", score)
		t.Fail()
	}
}

func TestGetMatch(t *testing.T) {
	cases := []struct {
		input    combination
		expected choice
	}{
		{combination{ROCK, DRAW}, ROCK},
		{combination{PAPER, LOSS}, ROCK},
		{combination{SCISSORS, WIN}, ROCK},
	}
	for _, c := range cases {
		choice := getMatch(c.input)
		if choice != c.expected {
			t.Error("test failed. ecpected", c.expected, "but got", choice)
			t.Fail()
		}
	}
}

func TestGetTotalScore2(t *testing.T) {
	input := []scorable{
		combination{ROCK, DRAW},
		combination{PAPER, LOSS},
		combination{SCISSORS, WIN},
	}
	expected := 12
	score := getTotalScore(input)
	if score != expected {
		t.Error("test failed. expected", expected, "but got", score)
		t.Fail()
	}
}
