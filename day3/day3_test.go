package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAsChoicepairInput(t *testing.T) {
	var input = []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg")
	expected := []rucksack{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
	}

	parsed := Parse(input)
	assert.Equal(t, expected, parsed, "parsed does not contain all expected values")
}

func TestFindMispack(t *testing.T) {
	rucksack := rucksack{"abc", "deb"}
	misPack := findMispack(rucksack)
	assert.Equal(t, 'b', misPack)
}

func TestFindMispacks(t *testing.T) {
	input := []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")
	parsed := Parse(input)
	commonItems := FindMispacks(parsed)
	assert.Equal(t, "pLPvts", commonItems)
}

func TestGetPriority(t *testing.T) {
	assert.Equal(t, 1, getPriority('a'))
	assert.Equal(t, 26, getPriority('z'))
	assert.Equal(t, 27, getPriority('A'))
	assert.Equal(t, 52, getPriority('Z'))
}

func TestGetPriorotySum(t *testing.T) {
	input := "pLPvts"
	sum := GetPrioritySum(input)
	assert.Equal(t, 157, sum)
}

func TestParseGroups(t *testing.T) {
	input := []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")
	firstExpected := group{
		[]rucksack{
			{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			{"PmmdzqPrV", "vPwwTWBwg"},
		},
	}
	parsed := ParseGroups(input)
	assert.Equal(t, firstExpected, parsed[0])
	assert.Equal(t, 2, len(parsed))
}

func TestMapItemTypes(t *testing.T) {
	input := group{
		[]rucksack{
			{"abc", "dez"},
			{"fgh", "zij"},
			{"zkl", "mno"},
		},
	}
	itemTypes := mapItemTypes(input)
	expected := map[rune]int{
		'a': 1,
		'b': 1,
		'c': 1,
		'd': 1,
		'e': 1,
		'f': 1,
		'g': 1,
		'h': 1,
		'i': 1,
		'j': 1,
		'k': 1,
		'l': 1,
		'm': 1,
		'n': 1,
		'o': 1,
		'z': 3,
	}
	assert.Equal(t, expected, itemTypes)
}

func TestFindBadgeType(t *testing.T) {
	input := []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")
	parsed := ParseGroups(input)
	badgeTypes := findBadgeTypes(parsed)
	assert.Equal(t, "rZ", badgeTypes)
}

func TestGetBadgePriorities(t *testing.T) {
	input := []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")
	parsed := ParseGroups(input)
	badgeTypes := findBadgeTypes(parsed)
	sum := GetPrioritySum(badgeTypes)
	assert.Equal(t, 70, sum)
}
