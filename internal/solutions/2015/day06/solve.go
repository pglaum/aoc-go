package main

import (
	"flag"
	"fmt"
	"os"
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

	lights := make([][]bool, 1000)
	for i := range 1000 {
		lights[i] = make([]bool, 1000)
	}

	for _, line := range lines {
		var instruction, state string
		var a1, a2, b1, b2 int
		if strings.Contains(line, "turn") {
			fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &instruction, &state, &a1, &a2, &b1, &b2)
		} else {
			fmt.Sscanf(line, "%s %d,%d through %d,%d", &instruction, &a1, &a2, &b1, &b2)
		}

		switch instruction {
		case "turn":
			switch state {
			case "on":
				for i := a1; i <= b1; i++ {
					for j := a2; j <= b2; j++ {
						lights[i][j] = true
					}
				}
			case "off":
				for i := a1; i <= b1; i++ {
					for j := a2; j <= b2; j++ {
						lights[i][j] = false
					}
				}
			}
		case "toggle":
			for i := a1; i <= b1; i++ {
				for j := a2; j <= b2; j++ {
					lights[i][j] = !lights[i][j]
				}
			}
		}
	}

	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				count += 1
			}
		}
	}
	return count
}

func part2(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)

	lights := make([][]int16, 1000)
	for i := range 1000 {
		lights[i] = make([]int16, 1000)
	}

	for _, line := range lines {
		var instruction, state string
		var a1, a2, b1, b2 int
		if strings.Contains(line, "turn") {
			fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &instruction, &state, &a1, &a2, &b1, &b2)
		} else {
			fmt.Sscanf(line, "%s %d,%d through %d,%d", &instruction, &a1, &a2, &b1, &b2)
		}

		switch instruction {
		case "turn":
			switch state {
			case "on":
				for i := a1; i <= b1; i++ {
					for j := a2; j <= b2; j++ {
						lights[i][j]++
					}
				}
			case "off":
				for i := a1; i <= b1; i++ {
					for j := a2; j <= b2; j++ {
						if lights[i][j] > 0 {
							lights[i][j]--
						}
					}
				}
			}
		case "toggle":
			for i := a1; i <= b1; i++ {
				for j := a2; j <= b2; j++ {
					lights[i][j] += 2
				}
			}
		}
	}

	for i := range lights {
		for j := range lights[i] {
			count += int(lights[i][j])
		}
	}
	return count
}
