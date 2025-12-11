package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gifts = `4x3x2
1x1x10
`

func TestSolve_2015_Day02(t *testing.T) {
	assert.Equal(t, 58+43, part1(gifts))
	assert.Equal(t, 34+14, part2(gifts))
}
