package day4

import (
	"log"
	"strconv"
	"strings"
)

// func Run(filename string) (int, int) {
// 	input := util.ReadFile(filename)
// 	parsed := Parse(input)
// 	return 0, 0
// }

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
	firstIndices := strings.Split(pair[0], "-")
	if len(firstIndices) != 2 {
		log.Fatal("could not split first indices")
	}
	secondIndices := strings.Split(pair[1], "-")
	if len(secondIndices) != 2 {
		log.Fatal("could not split second indices")
	}
	return elfPair{}
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
