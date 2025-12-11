package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var circuit = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
`

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 2, part1(circuit))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 2, part2(circuit))
	})
}
