package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var circuitSpec = `123 -> x
456 -> b
x AND b -> d
x OR b -> e
x LSHIFT 2 -> f
b RSHIFT 2 -> g
NOT x -> h
NOT b -> a
`

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 65079, part1(circuitSpec))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 456, part2(circuitSpec))
	})
}
