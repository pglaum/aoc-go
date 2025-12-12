package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var distances = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
`

func TestSolve_2015_Day01(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 605, part1(distances))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 982, part2(distances))
	})
}
