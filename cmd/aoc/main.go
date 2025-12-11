package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	initPtr = flag.Bool("init", false, "initialize solution template")
	yearPtr = flag.Int("year", 2025, "year of AoC puzzle")
	dayPtr  = flag.Int("day", 1, "day of AoC puzzle")
)

const pathTemplate = "./internal/solutions/%d/day%02d/solve.go"
const inputPathTemplate = "./inputs/%d/day%02d/%s.txt"

func main() {
	flag.Parse()

	inputPath := fmt.Sprintf(inputPathTemplate, *yearPtr, *dayPtr, "input")
	fmt.Println("input path:", inputPath)
	if fileInfo, err := os.Stat(inputPath); err != nil || fileInfo.IsDir() {
		panic("input file does not exist")
	}

	solutionPath := fmt.Sprintf(pathTemplate, *yearPtr, *dayPtr)
	fmt.Println("solution path:", solutionPath)
	if fileInfo, err := os.Stat(solutionPath); err != nil || fileInfo.IsDir() {
		panic("solution file does not exist")
	}

	cmd := exec.Command("go", "run", solutionPath, "-input", inputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
