package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var puzzle = `""
"abc"
"aaa\"aaa"
"\x27"
`

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 12, part1(puzzle))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 19, part2(puzzle))
	})
}
