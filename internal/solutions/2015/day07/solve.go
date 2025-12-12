package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
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

var circuit map[string]uint16

func getValue(s string) uint16 {
	num, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return uint16(num)
	}
	val, ok := circuit[s]
	if !ok {
		fmt.Printf("stack value not found: %q\n", s)
	}
	return val
}
func canGetValue(s string) bool {
	_, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return true
	}
	_, ok := circuit[s]
	return ok
}

func runWire(wiring []string, stack string) {
	switch len(wiring) {
	case 1:
		circuit[stack] = getValue(wiring[0])
	case 2:
		if wiring[0] == "NOT" {
			circuit[stack] = ^getValue(wiring[1])
		} else {
			fmt.Printf("unexpected wiring(2): %v -> %s\n", wiring, stack)
		}
	case 3:
		value1 := getValue(wiring[0])
		value2 := getValue(wiring[2])
		switch wiring[1] {
		case "AND":
			circuit[stack] = value1 & value2
		case "OR":
			circuit[stack] = value1 | value2
		case "LSHIFT":
			circuit[stack] = value1 << value2
		case "RSHIFT":
			circuit[stack] = value1 >> value2
		}
	}
}

func canRunWire(wiring []string) bool {
	switch len(wiring) {
	case 1:
		return canGetValue(wiring[0])
	case 2:
		return canGetValue(wiring[1])
	case 3:
		return canGetValue(wiring[0]) && canGetValue(wiring[2])
	}
	return false
}

func runSimulation(lines []string, skipBAssignment bool) {
	doneWires := []int{}
	for {
		for i := range lines {
			if slices.Contains(doneWires, i) {
				continue
			}
			line := lines[i]
			splitWiring := strings.Split(line, " -> ")
			wiring := strings.Fields(splitWiring[0])
			stack := splitWiring[1]

			if !canRunWire(wiring) {
				continue
			}

			if stack == "b" && len(wiring) == 1 && skipBAssignment {
				doneWires = append(doneWires, i)
				continue
			}

			runWire(wiring, stack)
			doneWires = append(doneWires, i)
		}
		if len(doneWires) >= len(lines) {
			break
		}
	}
}

func part1(puzzle string) int {
	lines := stringutils.SplitLines(puzzle, true)
	circuit = make(map[string]uint16)
	runSimulation(lines, false)
	return int(circuit["a"])
}

func part2(puzzle string) any {
	lines := stringutils.SplitLines(puzzle, true)
	circuit = make(map[string]uint16)

	runSimulation(lines, false)

	b := circuit["a"]
	circuit = make(map[string]uint16)
	circuit["b"] = b
	runSimulation(lines, true)

	return int(circuit["a"])
}
