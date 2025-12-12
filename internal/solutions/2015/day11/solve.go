package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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

func incrementPassword(current []byte) (next []byte) {
	carry := true
	for i := len(current) - 1; i >= 0; i-- {
		c := current[i]
		if carry {
			c += 1
		}
		if c > 'z' {
			c = 'a'
			carry = true
		} else {
			carry = false
		}
		next = append(next, c)
	}
	slices.Reverse(next)
	return next
}

func checkPassword(pw []byte) bool {
	hasAscendingLetters := false
	dupes := [][2]byte{}
	for i := range len(pw) - 1 {
		if i < (len(pw) - 2) {
			if pw[i]+2 == pw[i+1]+1 && pw[i]+2 == pw[i+2] {
				hasAscendingLetters = true
			}
		}
		if pw[i] == pw[i+1] {
			d := [2]byte{pw[i], pw[i+1]}
			if !slices.Contains(dupes, d) {
				dupes = append(dupes, d)
			}
		}
	}
	if !hasAscendingLetters {
		return false
	}
	if len(dupes) < 2 {
		return false
	}

	if slices.Contains(pw, 'i') || slices.Contains(pw, 'o') || slices.Contains(pw, 'l') {
		return false
	}

	return true
}

func part1(puzzle string) string {
	password := []byte(strings.TrimSpace(puzzle))
	for {
		password = incrementPassword(password)
		if checkPassword(password) {
			break
		}
	}
	return string(password)
}

func part2(puzzle string) string {
	password := []byte(strings.TrimSpace(puzzle))
	for {
		password = incrementPassword(password)
		if checkPassword(password) {
			break
		}
	}
	for {
		password = incrementPassword(password)
		if checkPassword(password) {
			break
		}
	}
	return string(password)
}
