package main

import (
	"flag"
	"fmt"
	"strconv"
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
	number, _ := strconv.Atoi(lines[0])
	elapsedParse := time.Since(start)

	start1 := time.Now()
	part1(number)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(number)
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func getNextRing(ring, edge, count, target int) (nextRing, nextEdge, nextCount int) {
	ring += 1
	count += edge*4 + 4
	edge += 2
	if count >= target {
		return ring, edge, count
	}
	return getNextRing(ring, edge, count, target)
}

func part1(number int) {
	ring, edge, count := getNextRing(1, 1, 1, number)
	fmt.Println("Ring:", ring, "Edge:", edge, "Count:", count)

	var vertical, horizontal int
	if number > count-edge {
		// bottom side
		vertical = ring - 1
		pos := ((edge + 1) / 2) - (count + 1 - number)
		if pos < 0 {
			horizontal = -pos
		} else {
			horizontal = pos
		}
	} else if number > count-(edge*2)+1 {
		// left side
		horizontal = ring - 1
		pos := ((edge + 1) / 2) - (count + 2 - edge - number)
		if pos < 0 {
			vertical = -pos
		} else {
			vertical = pos
		}
	} else if number > count-(edge*3)+2 {
		// top side
		vertical = ring - 1
		pos := ((edge + 1) / 2) - (count + 3 - (edge * 2) - number)
		if pos < 0 {
			horizontal = -pos
		} else {
			horizontal = pos
		}
	} else if number > count-(edge*4)+3 {
		// right side
		horizontal = ring - 1
		pos := ((edge + 1) / 2) - (count + 4 - (edge * 3) - number)
		if pos < 0 {
			vertical = -pos
		} else {
			vertical = pos
		}
	}
	println("Part 1:", vertical+horizontal)
}

func part2(number int) {
	num := 1

}
