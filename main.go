package main

import (
	"fmt"

	"github.com/nllptr/aoc2022/day1"
	"github.com/nllptr/aoc2022/day2"
	"github.com/nllptr/aoc2022/day3"
)

func main() {
	day1_1, day1_2 := day1.Run("day1/input.txt")
	fmt.Println("day 1, part 1:", day1_1)
	fmt.Println("day 1, part 2:", day1_2)
	day2_1, day2_2 := day2.Run("day2/input.txt")
	fmt.Println("day 2, part 1:", day2_1)
	fmt.Println("day 2, part 2:", day2_2)
	day3_1, day3_2 := day3.Run("day3/input.txt")
	fmt.Println("day 3, part 1:", day3_1)
	fmt.Println("day 3, part 2:", day3_2)
}
