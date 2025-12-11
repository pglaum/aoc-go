package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
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

func findSalt(puzzle, startsWith string) (salt int) {
	puzzle = strings.TrimSpace(puzzle)

	var saltStr string
	var hash []byte

	for {
		saltStr = strconv.Itoa(salt)
		hasher := md5.New()
		io.WriteString(hasher, puzzle+saltStr)
		hash = hasher.Sum(nil)

		if fmt.Sprintf("%x", hash)[:len(startsWith)] == startsWith {
			break
		}
		salt++
	}

	return salt
}

func part1(puzzle string) (salt int) {
	return findSalt(strings.TrimSpace(puzzle), "00000")
}

func part2(puzzle string) any {
	return findSalt(strings.TrimSpace(puzzle), "000000")
}
