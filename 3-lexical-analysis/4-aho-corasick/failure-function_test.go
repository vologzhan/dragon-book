package aho_corasick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAhoCorasickFailureFunction(t *testing.T) {
	for name, tt := range map[string]struct {
		patterns        []string
		failureFunction []int
	}{
		"": {
			patterns:        nil,
			failureFunction: nil,
		},
		"he_she_his_hers": {
			patterns:        []string{"he", "she", "his", "hers"},
			failureFunction: []int{0, 0, 0, 1, 2, 0, 3, 0, 3},
		},
		"aaa_abaaa_ababaaa": {
			patterns:        []string{"aaa", "abaaa", "ababaaa"},
			failureFunction: []int{0, 1, 2, 0, 1, 2, 3, 4, 5, 6, 7},
		},
		"all_fall_fatal_llama_lame": {
			patterns:        []string{"all", "fall", "fatal", "llama", "lame"},
			failureFunction: []int{0, 11, 12, 0, 1, 2, 3, 0, 1, 2, 0, 11, 16, 17, 1, 1, 0, 0},
		},
		"pipe_pet_item_temper_perpetual": {
			patterns:        []string{"pipe", "pet", "item", "temper", "perpetual"},
			failureFunction: []int{0, 7, 1, 5, 0, 11, 0, 11, 12, 13, 0, 0, 0, 1, 5, 17, 0, 1, 5, 6, 0, 0, 0},
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.failureFunction, getFailureFunction(tt.patterns))
		})
	}
}
