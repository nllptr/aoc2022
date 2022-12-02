package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse1Input(t *testing.T) {
	input := []byte("A Y\nB X\nC Z")
	expected := []choicePair{
		{ROCK, PAPER},
		{PAPER, ROCK},
		{SCISSORS, SCISSORS},
	}

	parsed := parse1(input)
	assert.Equal(t, expected, parsed, "parsed does not contain all expected values")
}
func TestParse2Input(t *testing.T) {
	input := []byte("A Y\nB X\nC Z")
	expected := []combination{
		{ROCK, DRAW},
		{PAPER, LOSS},
		{SCISSORS, WIN},
	}
	parsed := parse2(input)
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
		score := getScore(c.input)
		if score != c.expected {
			t.Error("test failed. expected", c.expected, "but got", score)
			t.Fail()
		}
	}
}

func TestGetTotalScore1(t *testing.T) {
	input := []choicePair{
		{ROCK, PAPER},
		{PAPER, ROCK},
		{SCISSORS, SCISSORS},
	}
	expected := 15
	score := getTotalScore1(input)
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
	input := []combination{
		{ROCK, DRAW},
		{PAPER, LOSS},
		{SCISSORS, WIN},
	}
	expected := 12
	score := getTotalScore2(input)
	if score != expected {
		t.Error("test failed. expected", expected, "but got", score)
		t.Fail()
	}
}
