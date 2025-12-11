package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var puzzle = "1"

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 2, part1(puzzle))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 2, part2(puzzle))
	})
}
