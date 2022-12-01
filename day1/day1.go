package day1

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/nllptr/aoc2022/util"
)

func Run(filename string) {
	input := util.ReadFile(filename)
	parsed := parse(input)
	max := max(parsed)
	fmt.Println("day 1, part 1:", max)
	maxThree := maxThree(parsed)
	fmt.Println("day 1, part 2:", maxThree)
}

func parse(input []byte) [][]int {
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

func max(input [][]int) int {
	max := 0
	for _, item := range input {
		sum := sum(item)
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxThree(input [][]int) int {
	sums := []int{}
	for _, item := range input {
		sums = append(sums, sum(item))
	}
	sort.Ints(sums)
	return sum(sums[len(sums)-3:])
}
