package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

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
	for _, line := range lines {
		unquoted, _ := strconv.Unquote(line)
		runeCount := utf8.RuneCountInString(unquoted)

		count += len(line) - runeCount
	}
	return count
}

func part2(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)
	for _, line := range lines {
		quoted := fmt.Sprintf("%q", line)
		count += len(quoted) - len(line)
	}
	return count
}
