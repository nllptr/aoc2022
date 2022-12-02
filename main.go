package main

import (
	"fmt"

	"github.com/nllptr/aoc2022/day1"
	"github.com/nllptr/aoc2022/day2"
)

func main() {
	var day1_1, day1_2 int
	day1_1, day1_2 = day1.Run("day1/input.txt")
	fmt.Println("day 1, part 1:", day1_1)
	fmt.Println("day 1, part 2:", day1_2)
	var day2_1, day2_2 int
	day2_1, day2_2 = day2.Run("day2/input.txt")
	fmt.Println("day 2, part 1:", day2_1)
	fmt.Println("day 2, part 2:", day2_2)
}
