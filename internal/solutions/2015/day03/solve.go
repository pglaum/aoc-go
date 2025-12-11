package main

import (
	"flag"
	"fmt"
	"image"
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

var nextPos = map[rune]image.Point{
	'>': image.Point{0, 1},
	'^': image.Point{1, 0},
	'v': image.Point{-1, 0},
	'<': image.Point{0, -1},
}

func part1(route string) int {
	route = stringutils.SplitLines(route, true)[0]
	pos := image.Point{0, 0}
	houses := []image.Point{pos}
	for _, step := range route {
		pos = pos.Add(nextPos[step])
		if !slices.Contains(houses, pos) {
			houses = append(houses, pos)
		}
	}
	return len(houses)
}

func part2(route string) int {
	route = stringutils.SplitLines(route, true)[0]
	posSanta := image.Point{0, 0}
	posRobot := image.Point{0, 0}
	houses := []image.Point{posSanta}
	for i := range len(route) / 2 {
		posSanta = posSanta.Add(nextPos[rune(route[i*2])])
		if !slices.Contains(houses, posSanta) {
			houses = append(houses, posSanta)
		}

		posRobot = posRobot.Add(nextPos[rune(route[i*2+1])])
		if !slices.Contains(houses, posRobot) {
			houses = append(houses, posRobot)
		}
	}
	return len(houses)
}
