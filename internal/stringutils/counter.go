package stringutils

func CountChars(input string) (count map[rune]int) {
	for _, c := range input {
		count[c] += 1
	}
	return count
}
