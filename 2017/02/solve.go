package main

import (
	"flag"
	"slices"
	"strconv"
	"strings"
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
	count := 0
	for _, line := range lines {
		nums := []int{}
		for snum := range strings.FieldsSeq(line) {
			num, _ := strconv.Atoi(snum)
			nums = append(nums, num)
		}
		count += slices.Max(nums) - slices.Min(nums)
	}
	println("Part 1:", count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		nums := []int{}
	Line:
		for snum := range strings.FieldsSeq(line) {
			num, _ := strconv.Atoi(snum)
			for _, existing := range nums {
				if num%existing == 0 {
					count += num / existing
					break Line
				} else if existing%num == 0 {
					count += existing / num
					break Line
				}
			}
			nums = append(nums, num)
		}
	}
	println("Part 2:", count)
}
