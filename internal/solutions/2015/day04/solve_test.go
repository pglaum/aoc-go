package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fiveZeros = map[string]int{
	"abcdef":  609043,
	"pqrstuv": 1048970,
}
var sixZeros = map[string]int{
	"abcdef":  6742839,
	"pqrstuv": 5714438,
}

func TestSolve_2015_Day04(t *testing.T) {
	for key, salt := range fiveZeros {
		t.Run("five zeros", func(t *testing.T) {
			assert.Equal(t, salt, part1(key))
		})
	}
	for key, salt := range sixZeros {
		t.Run("six zeros", func(t *testing.T) {
			assert.Equal(t, salt, part2(key))
		})
	}
	//assert.Equal(t, "1", part2(puzzle))
}
