package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/pglaum/aoc-go/internal/stringutils"
)

var (
	inputPath = flag.String("input", "", "path to input file")
	part      = flag.Int("part", 0, "part of the puzzle to solve (1 or 2). if 0, solve both parts.")
)

func main() {
	flag.Parse()

	if *inputPath == "" {
		panic("input file path is required")
	}

	puzzle, err := os.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("part 1: %v, took: %v\n", part1(string(puzzle)), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("part 2: %v, took: %v\n", part2(string(puzzle)), time.Since(start))
	}
}

func part1(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)
	sqFeet := 0
	for _, line := range lines {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)

		surface := 2*l*w + 2*l*h + 2*w*h

		sides := []int{l, w, h}
		slices.Sort(sides)

		sqFeet += surface + sides[0]*sides[1]
	}
	return sqFeet
}

func part2(puzzle string) any {
	lines := stringutils.SplitLines(puzzle, true)
	ribbon := 0
	for _, line := range lines {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)

		sides := []int{l, w, h}
		slices.Sort(sides)

		ribbon += 2*sides[0] + 2*sides[1] + l*w*h
	}
	return ribbon
}
