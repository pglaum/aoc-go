package util

import (
	"os"
	"strings"
)

func ReadInputLines(filename string, trim bool) []string {
	data, _ := os.ReadFile(filename)

	input := strings.ReplaceAll(string(data), "\r\n", "\n")
	if trim {
		input = strings.TrimSpace(input)
	}
	lines := strings.Split(input, "\n")
	if !trim && lines[len(lines)-1] == "" {
		// last line is a newline
		lines = lines[:len(lines)-1]
	}

	return lines
}
