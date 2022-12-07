package main

import (
	"fmt"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := strings.TrimSpace(shared.ReadFileContents("input.txt"))
	solution := solve(contents)
	fmt.Printf("Solution: %v\n", solution)
}

func solve(contents string) int {
	buffer := contents[:14]
	if allDifferent(buffer) {
		return 0
	}
	for i, c := range contents[14:] {
		buffer = buffer[1:] + string(c)
		if allDifferent(buffer) {
			return i + 15
		}
	}
	return -1
}

func allDifferent(buffer string) bool {
	check := make(map[rune]bool, 0)
	for _, c := range buffer {
		if check[c] {
			return false
		}
		check[c] = true
	}
	return true
}
