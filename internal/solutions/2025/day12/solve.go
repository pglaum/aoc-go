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
		fmt.Printf("there is no part 2 for this puzzle")
	}
}

func part1(puzzle string) (count int) {
	// naive approach: just check the area size
	regions := strings.Split(strings.ReplaceAll(puzzle, "\r\n", "\n"), "\n\n")
	trees := strings.SplitSeq(strings.TrimSpace(regions[len(regions)-1]), "\n")
	for tree := range trees {
		var x, y, a, b, c, d, e, f int
		fmt.Sscanf(tree, "%dx%d: %d %d %d %d %d %d", &x, &y, &a, &b, &c, &d, &e, &f)
		naivePresentArea := (a + b + c + d + e + f) * 3 * 3
		treeArea := x * y
		if (float64(naivePresentArea) * 0.88) < float64(treeArea) {
			count++
		}
	}

	return count
}
