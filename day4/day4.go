package day4

import (
	"log"
	"strconv"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) (int, int) {
	input := util.ReadFile(filename)
	parsed := Parse(input)
	fullyContained := countContainsTheOther(parsed)
	overlaps := countHasOverlap(parsed)
	return fullyContained, overlaps
}

type elfPair struct {
	first  []int
	second []int
}

func Parse(input []byte) []elfPair {
	asString := string(input)
	rows := strings.Split(asString, "\n")
	parsed := []elfPair{}
	for _, row := range rows {
		elfPair := asElfPair(row)
		parsed = append(parsed, elfPair)
	}
	return parsed
}

func asElfPair(row string) elfPair {
	pair := strings.Split(row, ",")
	if len(pair) != 2 {
		log.Fatal("could not split row ", pair)
	}
	firstStart, firstEnd := splitToIndices(pair[0])
	firstString := indicesToSlice(firstStart, firstEnd)
	secondStart, secondEnd := splitToIndices(pair[1])
	secondString := indicesToSlice(secondStart, secondEnd)
	return elfPair{firstString, secondString}
}

func splitToIndices(indicesStr string) (int, int) {
	stringIndices := strings.Split(indicesStr, "-")
	if len(stringIndices) != 2 {
		log.Fatal("could not split first indices")
	}
	return toInts(stringIndices)
}

func toInts(stringIndices []string) (int, int) {
	firstIndex, err := strconv.Atoi(stringIndices[0])
	if err != nil {
		log.Fatal("could not parse as int", firstIndex)
	}
	secondIndex, err := strconv.Atoi(stringIndices[1])
	if err != nil {
		log.Fatal("could not parse as int", secondIndex)
	}
	return firstIndex, secondIndex
}

func indicesToSlice(start, end int) []int {
	result := []int{}
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func oneContainsTheOther(elfPair elfPair) bool {
	if containsAll(elfPair.first, elfPair.second) {
		return true
	}
	if containsAll(elfPair.second, elfPair.first) {
		return true
	}
	return false
}

func containsAll(first, second []int) bool {
	for _, v := range second {
		if !contains(first, v) {
			return false
		}
	}
	return true
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func countContainsTheOther(elfPairs []elfPair) int {
	count := 0
	for _, e := range elfPairs {
		if oneContainsTheOther(e) {
			count = count + 1
		}
	}
	return count
}

func countHasOverlap(elfPairs []elfPair) int {
	count := 0
	for _, e := range elfPairs {
		if hasOverlap(e) {
			count = count + 1
		}
	}
	return count
}

func hasOverlap(elfPair elfPair) bool {
	for _, v := range elfPair.second {
		if contains(elfPair.first, v) {
			return true
		}
	}
	return false
}
