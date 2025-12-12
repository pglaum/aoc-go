package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var passwords = map[string]string{
	"abcdefgh": "abcdffaa",
	"ghijklmn": "ghjaabcc",
}

func TestSolve_2015_Day11(t *testing.T) {
	for current, next := range passwords {
		t.Run("part 1", func(t *testing.T) {
			assert.Equal(t, next, part1(current))
		})
	}
}
