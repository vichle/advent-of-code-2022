package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("./input.txt")

	rows := strings.Split(contents, "\n")

	current := 0
	totals := make([]int, 0)
	for _, row := range rows {
		if row == "" {
			totals = append(totals, current)
			current = 0
		}
		calories, _ := strconv.Atoi(row)
		current += calories
	}
	sort.Ints(totals)
	l := len(totals)
	fmt.Printf("Max calories: %v\n", totals[l-1])
	fmt.Printf("Sum top three calories: %v\n", totals[l-1]+totals[l-2]+totals[l-3])
}
