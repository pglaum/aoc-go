package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var instructions = `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500
`

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 998996, part1(instructions))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 1001996, part2(instructions))
	})
}
