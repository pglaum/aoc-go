package stringutils

import "strings"

func SplitLines(input string, trim bool) (lines []string) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	if trim {
		input = strings.TrimSpace(input)
	}
	lines = strings.Split(input, "\n")
	if !trim {
		lines = lines[:len(lines)-1]
	}
	return lines
}
