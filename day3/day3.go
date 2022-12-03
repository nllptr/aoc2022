package day3

import (
	"errors"
	"log"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) (int, int) {
	input := util.ReadFile(filename)
	parsed := Parse(input)
	misPacks := FindMispacks(parsed)
	sum := GetPrioritySum(misPacks)
	parsedGroups := ParseGroups(input)
	badgeTypes := findBadgeTypes(parsedGroups)
	badgePriorities := GetPrioritySum(badgeTypes)
	return sum, badgePriorities
}

type rucksack struct {
	compOne string
	compTwo string
}

type group struct {
	rucksacks []rucksack
}

func Parse(input []byte) []rucksack {
	asString := string(input)
	rows := strings.Split(asString, "\n")
	parsed := []rucksack{}
	for _, row := range rows {
		if len(row)%2 != 0 {
			log.Fatal("row can't be split in half", row)
		}
		middle := len(row) / 2
		rucksack := rucksack{row[:middle], row[middle:]}
		parsed = append(parsed, rucksack)
	}
	return parsed
}

func findMispack(rucksack rucksack) rune {
	var common rune
	for _, f := range rucksack.compOne {
		for _, s := range rucksack.compTwo {
			if f == s {
				common = f
			}
		}
	}
	if common == 0 {
		log.Fatal("no common ocurrence found")
	}
	return common
}

func FindMispacks(rucksacks []rucksack) string {
	var sb strings.Builder
	for _, rucksack := range rucksacks {
		misPack := findMispack(rucksack)
		sb.WriteRune(misPack)
	}
	return sb.String()
}

func getPriority(r rune) int {
	intVal := int(r)
	if intVal >= 97 && intVal <= 122 {
		intVal = intVal - 96
	} else if intVal >= 65 && intVal <= 90 {
		intVal = intVal - 38
	} else {
		log.Fatal("invalid character", r)
	}
	return intVal
}

func GetPrioritySum(misPacks string) int {
	sum := 0
	for _, m := range misPacks {
		sum = sum + getPriority(m)
	}
	return sum
}

func ParseGroups(input []byte) []group {
	rucksacks := Parse(input)
	groups := []group{}
	parsed, _ := appendGroup(groups, rucksacks)
	return parsed
}

func appendGroup(groups []group, rucksacks []rucksack) ([]group, []rucksack) {
	if len(rucksacks) == 0 {
		return groups, nil
	}
	if len(rucksacks) <= 3 {
		return append(groups, group{rucksacks}), nil
	}
	firstThree := rucksacks[:3]
	remaining := rucksacks[3:]
	groups = append(groups, group{firstThree})
	return appendGroup(groups, remaining)
}

func findBadgeTypes(groups []group) string {
	var sb strings.Builder
	for _, g := range groups {
		badgeType, err := findBadgeType(g)
		if err != nil {
			log.Fatal(err)
		}
		sb.WriteRune(badgeType)
	}
	return sb.String()
}

func findBadgeType(group group) (rune, error) {
	itemMap := mapItemTypes(group)
	for item, val := range itemMap {
		if val == 3 {
			return item, nil
		}
	}
	return 0, errors.New("could not find a badge type")
}

func mapItemTypes(group group) map[rune]int {
	groupMap := map[rune]int{}
	for _, rucksack := range group.rucksacks {
		rucksackMap := map[rune]bool{}
		for _, item := range rucksack.compOne + rucksack.compTwo {
			if !rucksackMap[item] {
				rucksackMap[item] = true
			}
		}
		for item, _ := range rucksackMap {
			groupMap[item] = groupMap[item] + 1
		}
	}
	return groupMap
}
