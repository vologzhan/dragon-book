package aho_corasick

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAhoCorasickFailureFunction(t *testing.T) {
	for input, expected := range map[string][]int{
		"ababaa":    {0, 0, 1, 2, 3, 1},
		"abababaab": {0, 0, 1, 2, 3, 4, 5, 1, 2},
		"aaaaaa":    {0, 1, 2, 3, 4, 5},
		"abbaabb":   {0, 0, 0, 1, 1, 2, 3},
	} {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, getFailureFunction(input))
		})
	}
}

func TestAhoCorasickSearch(t *testing.T) {
	for input, expected := range map[string]bool{
		"ababaa":     true,
		"aabbababaa": true,
		"abababaab":  true,
		"ababa":      false,
		"abababbaa":  false,
	} {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, search(input, "ababaa"))
		})
	}
}

func TestFibonacciString(t *testing.T) {
	for input, expected := range map[string]string{
		"1": "b",
		"2": "a",
		"3": "ab",
		"4": "aba",
		"5": "abaab",
		"6": "abaababa",
		"7": "abaababaabaab",
		"8": "abaababaabaababaababa",
	} {
		t.Run(input, func(t *testing.T) {
			s, _ := strconv.Atoi(input)
			assert.Equal(t, expected, getFibonacciString(s))
		})
	}
}
