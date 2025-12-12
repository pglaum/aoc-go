package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/pglaum/aoc-go/internal/graphutils"
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

func bfs(G *graphutils.Graph, from string, visited []string, currLen int, longest bool) (best int) {
	if longest {
		best = 0
	} else {
		best = int(^uint(16) >> 1)
	}
	notFound := true
	for _, edge := range G.FindEdges(from) {
		if slices.Contains(visited, edge.To) {
			continue
		}
		notFound = false
		length := bfs(G, edge.To, append(visited, edge.To), currLen+edge.Weight, longest)
		if longest {
			best = max(best, length)
		} else {
			best = min(best, length)
		}
	}
	if notFound {
		return currLen
	}

	return best
}

func part1(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)
	g := graphutils.NewGraph()
	for _, line := range lines {
		var from, to string
		var dist int
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist)
		g.AddEdge(graphutils.Edge{From: from, To: to, Weight: dist})
	}

	for _, node := range g.Nodes {
		shortestForNode := bfs(g, node.ID, []string{node.ID}, 0, false)
		if count == 0 || shortestForNode < count {
			count = shortestForNode
		}
	}
	return count
}

func part2(puzzle string) (count int) {
	lines := stringutils.SplitLines(puzzle, true)
	g := graphutils.NewGraph()
	for _, line := range lines {
		var from, to string
		var dist int
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist)
		g.AddEdge(graphutils.Edge{From: from, To: to, Weight: dist})
	}

	for _, node := range g.Nodes {
		longestForNode := bfs(g, node.ID, []string{node.ID}, 0, true)
		if longestForNode > count {
			count = longestForNode
		}
	}
	return count
}
