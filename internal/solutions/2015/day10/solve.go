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

func lookAndSay(input []byte) (output []byte) {
	i := 0
	for i < len(input) {
		count := 1
		for j := i + 1; j < len(input); j++ {
			if input[j] != input[i] {
				break
			}
			count++
		}
		output = fmt.Appendf(output, "%d%c", count, input[i])
		i += count
	}
	return
}

func part1(puzzle string) (count int) {
	say := []byte(strings.TrimSpace(puzzle))
	for range 40 {
		say = lookAndSay(say)
	}
	return len(say)
}

func part2(puzzle string) (count int) {
	say := []byte(strings.TrimSpace(puzzle))
	for range 50 {
		say = lookAndSay(say)
	}
	return len(say)
}
