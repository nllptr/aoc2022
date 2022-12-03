package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) (int, int) {
	input := util.ReadFile(filename)
	parsed := Parse(input)
	max := Max(parsed)
	maxThree := MaxThree(parsed)
	return max, maxThree
}

func Parse(input []byte) [][]int {
	asString := string(input)
	stringGroups := strings.Split(asString, "\n\n")
	parsed := [][]int{}
	for _, group := range stringGroups {
		group := strings.Split(group, "\n")
		intGroup := []int{}
		for _, stringValue := range group {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				log.Fatal("could not convert to int:", stringValue)
			}
			intGroup = append(intGroup, value)
		}
		parsed = append(parsed, intGroup)
	}
	return parsed
}

func sum(items []int) int {
	sum := 0
	for _, val := range items {
		sum = sum + val
	}
	return sum
}

func Max(input [][]int) int {
	max := 0
	for _, item := range input {
		sum := sum(item)
		if sum > max {
			max = sum
		}
	}
	return max
}

func MaxThree(input [][]int) int {
	sums := []int{}
	for _, item := range input {
		sums = append(sums, sum(item))
	}
	sort.Ints(sums)
	return sum(sums[len(sums)-3:])
}
