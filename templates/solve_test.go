package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var puzzle = "1"

func TestSolve_2015_Day01(t *testing.T) {
	assert.Equal(t, "1", part1(puzzle))
	assert.Equal(t, "1", part2(puzzle))
}
