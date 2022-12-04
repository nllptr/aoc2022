package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nllptr/aoc2022/day1"
	"github.com/nllptr/aoc2022/day2"
	"github.com/nllptr/aoc2022/day3"
	"github.com/nllptr/aoc2022/day4"
	"github.com/stretchr/testify/assert"
)

// These tests are mainly to make sure refactorings do not introduce any regressions.
func TestAll(t *testing.T) {
	day1_1, day1_2 := day1.Run("day1/input.txt")
	assert.Equal(t, 70720, day1_1)
	assert.Equal(t, 207148, day1_2)

	day2_1, day2_2 := day2.Run("day2/input.txt")
	assert.Equal(t, 17189, day2_1)
	assert.Equal(t, 13490, day2_2)

	day3_1, day3_2 := day3.Run("day3/input.txt")
	assert.Equal(t, 8139, day3_1)
	assert.Equal(t, 2668, day3_2)

	day4_1, day4_2 := day4.Run("day4/input.txt")
	assert.Equal(t, 503, day4_1)
	assert.Equal(t, 827, day4_2)
}

func BenchmarkDay1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		day1.Run("day1/input.txt")
	}
}

func BenchmarkDay2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		day2.Run("day2/input.txt")
	}
}

func BenchmarkDay3(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		day3.Run("day3/input.txt")
	}
}

func BenchmarkDay4(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		day4.Run("day4/input.txt")
	}
}
