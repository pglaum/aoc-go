package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var samples1 = map[string]int{
	"(())":    0,
	"()()":    0,
	"(((":     3,
	"(()(()(": 3,
	"))(((((": 3,
	"())":     -1,
	"))(":     -1,
	")))":     -3,
	")())())": -3,
}
var samples2 = map[string]int{
	")":     1,
	"()())": 5,
}

func TestSolve_2015_Day01(t *testing.T) {
	num := 0
	for input, result := range samples1 {
		t.Run(fmt.Sprintf("part 1 run %d", num), func(t *testing.T) {
			assert.Equal(t, result, part1(input))
		})
		num += 1
	}
	num = 0
	for input, result := range samples2 {
		t.Run(fmt.Sprintf("part 2 run %d", num), func(t *testing.T) {
			assert.Equal(t, result, part2(input))
		})
		num += 1
	}
}
