package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var input = []byte("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8")
	expected := []elfPair{
		{
			[]int{2, 3, 4},
			[]int{6, 7, 8},
		},
		{
			[]int{2, 3},
			[]int{4, 5},
		},
		{
			[]int{5, 6, 7},
			[]int{7, 8, 9},
		},
		{
			[]int{2, 3, 4, 5, 6, 7, 8},
			[]int{3, 4, 5, 6, 7},
		},
		{
			[]int{6},
			[]int{4, 5, 6},
		},
		{
			[]int{2, 3, 4, 5, 6},
			[]int{4, 5, 6, 7, 8},
		},
	}
	parsed := Parse(input)
	assert.Equal(t, expected, parsed, "parsed does not contain all expected values")
}

func TestToInts(t *testing.T) {
	input := []string{"123", "456"}
	first, second := toInts(input)
	assert.Equal(t, 123, first)
	assert.Equal(t, 456, second)
}

func TestSplitToIndices(t *testing.T) {
	input := "123-456"
	first, second := splitToIndices(input)
	assert.Equal(t, 123, first)
	assert.Equal(t, 456, second)
}

func TestIndicesToString(t *testing.T) {
	str := indicesToSlice(4, 6)
	assert.Equal(t, []int{4, 5, 6}, str)

	str = indicesToSlice(6, 6)
	assert.Equal(t, []int{6}, str)

	str = indicesToSlice(123, 125)
	assert.Equal(t, []int{123, 124, 125}, str)
}

func TestOneContainsTheOther(t *testing.T) {
	cases := []struct {
		input    elfPair
		expected bool
	}{
		{
			elfPair{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
			},
			true,
		},
		{
			elfPair{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			false,
		},
		{
			elfPair{
				[]int{6},
				[]int{1, 2, 3, 4, 5, 6},
			},
			true,
		},
	}
	for _, c := range cases {
		result := oneContainsTheOther(c.input)
		assert.Equal(t, c.expected, result)
	}
}

func TestContains(t *testing.T) {
	input := []int{1, 2, 3}
	assert.True(t, contains(input, 2))
	assert.False(t, contains(input, 5))
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		first    []int
		second   []int
		expected bool
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 5, 3},
			false,
		},
		{
			[]int{1, 2, 3},
			[]int{3},
			true,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, containsAll(c.first, c.second))
	}
}

func TestCountContainsTheOther(t *testing.T) {
	input := []byte("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8")
	parsed := Parse(input)
	assert.Equal(t, 2, countContainsTheOther(parsed))
}

func TestCountHasOverlap(t *testing.T) {
	input := []byte("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8")
	parsed := Parse(input)
	assert.Equal(t, 4, countHasOverlap(parsed))
}
