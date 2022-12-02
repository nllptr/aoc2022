package main

import (
	"testing"

	"github.com/nllptr/aoc2022/day1"
	"github.com/nllptr/aoc2022/day2"
)

// These tests are mainly to make sure refactorings do not introduce any regressions.
func TestAll(t *testing.T) {
	day1_1, day1_2 := day1.Run("day1/input.txt")
	if day1_1 != 70720 {
		t.Error("day1_1 is broken")
		t.Fail()
	}
	if day1_2 != 207148 {
		t.Error("day1_2 is broken")
		t.Fail()
	}

	day2_1, day2_2 := day2.Run("day2/input.txt")
	if day2_1 != 17189 {
		t.Error("day2_1 is broken")
		t.Fail()
	}
	if day2_2 != 13490 {
		t.Error("day2_2 is broken")
		t.Fail()
	}
}
