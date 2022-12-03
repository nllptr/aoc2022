package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	input := []byte("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000")
	expected := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	parsed := Parse(input)
	assert.Equal(t, parsed, expected, "parsed does not contain all expected values")
}

func TestSumCalories(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{1000, 2000, 3000}, 6000},
		{[]int{4000}, 4000},
		{[]int{5000, 6000}, 11000},
		{[]int{7000, 8000, 9000}, 24000},
		{[]int{10000}, 10000},
	}
	for _, c := range cases {
		sum := sum(c.input)
		if sum != c.expected {
			t.Error("test failed. expected", c.expected, "but got", sum)
			t.Fail()
		}
	}
}

func TestGetMax(t *testing.T) {
	input := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
	expected := 24000
	max := Max(input)
	if max != expected {
		t.Error("test failed. expected", expected, "but got", max)
		t.Fail()
	}
}

func TestGetMaxThree(t *testing.T) {
	input := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
	expected := 45000
	maxThree := MaxThree(input)
	if maxThree != expected {
		t.Error("expected sum to be", expected, "but got", maxThree)
		t.Fail()
	}
}
