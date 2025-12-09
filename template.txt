package main

import (
	"flag"
	"time"

	"github.com/pglaum/aoc-go/util"
)

var sample bool

func main() {
	start := time.Now()
	flag.BoolVar(&sample, "sample", false, "use sample input")
	flag.Parse()

	filename := "input.txt"
	if sample {
		filename = "sample.txt"
	}
	lines := util.ReadInputLines(filename, true)
	elapsedParse := time.Since(start)

	start1 := time.Now()
	part1(lines)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(lines)
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func part1(lines []string) {
}

func part2(lines []string) {
}
