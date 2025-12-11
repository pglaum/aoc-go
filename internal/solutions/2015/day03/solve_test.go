package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sample1 = map[string]int{
	">":          2,
	"^>v<":       4,
	"^v^v^v^v^v": 2,
}

var sample2 = map[string]int{
	"^v":         3,
	"^>v<":       3,
	"^v^v^v^v^v": 11,
}

func TestSolve_2015_Day03(t *testing.T) {
	for route, houses := range sample1 {
		t.Run("part 1", func(t *testing.T) {
			assert.Equal(t, houses, part1(route))
		})
	}
	for route, houses := range sample2 {
		t.Run("part 2", func(t *testing.T) {
			assert.Equal(t, houses, part2(route))
		})
	}
}
