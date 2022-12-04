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
