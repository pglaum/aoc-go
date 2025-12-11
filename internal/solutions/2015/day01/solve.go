package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func part1(puzzle string) any {
	puzzle = strings.TrimSpace(puzzle)
	floor := 0
	for _, char := range puzzle {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			println("unexpected char:", char)
		}
	}
	return floor
}

func part2(puzzle string) any {
	puzzle = strings.TrimSpace(puzzle)
	floor := 0
	for i, char := range puzzle {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			println("unexpected char:", char)
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}
