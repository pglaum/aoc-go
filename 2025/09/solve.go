package main

import (
	"flag"
	"fmt"
	"image"
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

	points := []image.Point{}
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, image.Point{x, y})
	}
	elapsedParse := time.Since(start)

	start1 := time.Now()
	part1(points)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(points)
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func part1(points []image.Point) {
	largestArea := 0
	rects := []image.Rectangle{}
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			rects = append(rects, image.Rectangle{p1, p2}.Canon())
		}
	}
	for _, r := range rects {
		// calculate the correct size including borders
		r.Max = r.Max.Add(image.Point{1, 1})
		area := r.Dx() * r.Dy()
		largestArea = max(largestArea, area)
	}
	println("Part 1:", largestArea)
}

func part2(points []image.Point) {
	largestArea := 0
	var rects, lines []image.Rectangle
	for i, p1 := range points {
		for j, p2 := range points[i+1:] {
			if j == 0 {
				lines = append(lines, image.Rectangle{p1, p2}.Canon())
			}
			rects = append(rects, image.Rectangle{p1, p2}.Canon())
		}
	}
	lines = append(lines, image.Rectangle{points[len(points)-1], points[0]}.Canon())

Loop:
	for _, r := range rects {
		// calculate the correct size including borders
		r.Max = r.Max.Add(image.Point{1, 1})
		area := r.Dx() * r.Dy()

		for _, l := range lines {
			l.Max = l.Max.Add(image.Point{1, 1})
			if l.Overlaps(r.Inset(1)) {
				continue Loop
			}
		}
		largestArea = max(largestArea, area)
	}
	println("Part 2:", largestArea)
}
