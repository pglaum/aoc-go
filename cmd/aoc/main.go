package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/pglaum/aoc-go/internal/utils"
)

var (
	initPtr = flag.Bool("init", false, "initialize solution template")
	testPtr = flag.Bool("test", false, "run tests for a day")
	yearPtr = flag.Int("year", 2025, "year of AoC puzzle")
	dayPtr  = flag.Int("day", 1, "day of AoC puzzle")
)

const pathTemplate = "./internal/solutions/%d/day%02d/solve.go"
const inputPathTemplate = "./inputs/%d/day%02d/%s.txt"

var solveFiles = [2]string{
	"solve.go",
	"solve_test.go",
}

func main() {
	flag.Parse()

	if *initPtr {
		if err := initSolutionTemplate(*yearPtr, *dayPtr); err != nil {
			panic(err)
		}
		fmt.Println("initialized solution template: " + fmt.Sprintf(pathTemplate, *yearPtr, *dayPtr))
		return
	}

	if *testPtr {
		testPath := fmt.Sprintf("./internal/solutions/%d/day%02d", *yearPtr, *dayPtr)

		cmd := exec.Command("go", "test", "-v", testPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		return
	}

	inputPath, err := utils.FetchInput(*yearPtr, *dayPtr)
	if err != nil {
		panic("input file does not exist")
	}
	fmt.Println("input path:", inputPath)

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

func initSolutionTemplate(year, day int) error {
	solutionDir := fmt.Sprintf("./internal/solutions/%d/day%02d", year, day)
	if err := os.MkdirAll(solutionDir, os.ModePerm); err != nil {
		return err
	}

	for _, solveFile := range solveFiles {
		solutionPath := fmt.Sprintf("%s/%s", solutionDir, solveFile)
		templatePath := fmt.Sprintf("./templates/%s", solveFile)

		input, err := os.ReadFile(templatePath)
		if err != nil {
			return err
		}

		if fileInfo, err := os.Stat(solutionPath); err == nil && !fileInfo.IsDir() {
			return fmt.Errorf("file already exists: %s", solutionPath)
		}

		if err := os.WriteFile(solutionPath, input, 0644); err != nil {
			return err
		}
	}

	return nil
}
