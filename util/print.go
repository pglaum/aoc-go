package util

import (
	"fmt"
	"time"
)

func PrintElapsed(parse, part1, part2, full time.Duration) {
	fmt.Println()
	fmt.Println("time taken")
	fmt.Printf("  parse: %s\n", parse)
	fmt.Printf("  part1: %s\n", part1)
	fmt.Printf("  part2: %s\n", part2)
	fmt.Printf("  full: %s\n", full)
}
