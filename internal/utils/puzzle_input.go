package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pglaum/aoc-go/pkg/aoc"
)

const (
	cachePath     = "./inputs/%d/day%02d/%s.txt"
	inputFileName = "input"
)

func FetchInput(year, day int) (string, error) {
	fileName := fmt.Sprintf(cachePath, year, day, inputFileName)
	if fileInfo, err := os.Stat(fileName); err == nil && !fileInfo.IsDir() {
		return fileName, nil
	}

	client, err := aoc.NewAocClient()
	if err != nil {
		return "", err
	}

	input, err := client.FetchInput(year, day)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		return "", err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	_, err = f.WriteString(input)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
