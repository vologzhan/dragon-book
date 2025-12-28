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
