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
	part1(lines[0])
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(lines[0])
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func part1(line string) {
	sum := 0
	for i := range line {
		if i == len(line)-1 {
			if line[i] == line[0] {
				sum += int(line[i] - '0')
			}
		} else {
			if line[i] == line[i+1] {
				sum += int(line[i] - '0')
			}
		}
	}
	println("Part 1:", sum)
}

func part2(line string) {
	sum := 0
	for i := range line[:len(line)/2] {
		if line[i] == line[i+len(line)/2] {
			sum += int(line[i]-'0') * 2
		}
	}
	println("Part 2:", sum)
}
