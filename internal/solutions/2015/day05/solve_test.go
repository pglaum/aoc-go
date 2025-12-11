package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sample1 = `ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb
`

var sample2 = `qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy
`

func TestSolve_2015_Day05(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 2, part1(sample1))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 2, part2(sample2))
	})
}
