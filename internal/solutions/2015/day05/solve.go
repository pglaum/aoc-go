package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
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
	badStrings := []string{"ab", "cd", "pq", "xy"}
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

Line:
	for _, line := range lines {
		for _, bad := range badStrings {
			if strings.Contains(line, bad) {
				continue Line
			}
		}

		vowelCount := 0
		hasDupe := false
		for i, char := range line {
			if slices.Contains(vowels, char) {
				vowelCount++
			}
			if !hasDupe && i > 0 {
				if line[i-1] == line[i] {
					hasDupe = true
				}
			}
		}

		if vowelCount >= 3 && hasDupe {
			count++
		}
	}

	return count
}

func part2(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)

	for _, line := range lines {
		hasDupe := false
		hasMirror := false

		for i := range line {
			if i > 0 && !hasDupe {
				testDupe := line[i-1 : i+1]
				if strings.Contains(line[i+1:], testDupe) {
					hasDupe = true
				}
			}
			if i > 1 && !hasMirror {
				if line[i-2] == line[i] {
					hasMirror = true
				}
			}
		}

		if hasDupe && hasMirror {
			count++
		}
	}

	return count
}
