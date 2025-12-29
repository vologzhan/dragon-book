package aho_corasick

import (
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
