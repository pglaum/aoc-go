package main

import (
	"flag"
	"slices"
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

	var edges = [][2]string{}
	for _, line := range lines {
		fields := strings.Fields(line)
		from := fields[0]
		from = from[:len(from)-1]
		for _, field := range fields[1:] {
			edges = append(edges, [2]string{from, field})
		}
	}

	elapsedParse := time.Since(start)

	start1 := time.Now()
	part1(edges)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(edges)
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func getOut(edges [][2]string, from string, path []string) int {
	count := 0
	possibilities := []string{}
	for _, edge := range edges {
		if edge[0] == from {
			if slices.Contains(path, edge[1]) {
				// no cycle
				continue
			}
			possibilities = append(possibilities, edge[1])
		}
	}

	for _, to := range possibilities {
		if to == "out" {
			count += 1
			continue
		}
		count += getOut(edges, to, append(path, to))
	}
	return count
}

func part1(edges [][2]string) {
	count := getOut(edges, "you", []string{"you"})
	println("Part 1:", count)
}

type State struct {
	nodeID string
	hasDac bool
	hasFft bool
}

func part2(edges [][2]string) {
	g := util.NewGraph()
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}
	startNode, _ := g.GetNode("svr")
	count := dfs(startNode, "out", false, false, make(map[string]bool), make(map[State]int))
	println("Part 1:", count)
}

func dfs(node *util.Node, target string, hasDac, hasFft bool, visited map[string]bool, cache map[State]int) int {
	if node.ID == "dac" {
		hasDac = true
	}
	if node.ID == "fft" {
		hasFft = true
	}
	state := State{node.ID, hasDac, hasFft}

	if count, found := cache[state]; found {
		return count
	}

	if node.ID == target {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	visited[node.ID] = true

	count := 0
	for _, neighbor := range node.Adjacent {
		if !visited[neighbor.ID] {
			count += dfs(neighbor, target, hasDac, hasFft, visited, cache)
		}
	}

	delete(visited, node.ID)

	cache[state] = count
	return count
}
