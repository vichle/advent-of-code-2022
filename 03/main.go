package main

import (
	"fmt"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	rows := strings.Split(contents, "\n")
	duplicates := make([]int, 0)
	for _, row := range rows {
		splitSacks := strings.Split(row, "")
		firstRuckSackItems := strings.Join(splitSacks[:len(splitSacks)/2], "")
		secondRuckSackItems := strings.Join(splitSacks[len(splitSacks)/2:], "")
		set := make(map[int]bool, 0)
		for _, c := range firstRuckSackItems {
			set[int(c)] = true
		}
		for _, c := range secondRuckSackItems {
			if set[int(c)] {
				duplicates = append(duplicates, int(c))
				break
			}
		}
	}

	prioSum := 0
	for _, n := range duplicates {
		prioSum += getPrio(n)
	}
	fmt.Printf("Priority sum: %v\n", prioSum)

	var sharedItems map[int]bool
	badgePrioSum := 0
	for i, row := range rows {
		itemSet := make(map[int]bool)
		for _, c := range row {
			itemSet[int(c)] = true
		}
		if i%3 == 0 {
			sharedItems = itemSet
		} else {
			for item, _ := range sharedItems {
				if !itemSet[item] {
					delete(sharedItems, item)
				}
			}
		}
		if i%3 == 2 {
			keys := make([]int, 0, len(sharedItems))
			for k := range sharedItems {
				keys = append(keys, k)
			}

			badgePrioSum += getPrio(keys[0])
		}
	}
	fmt.Printf("Badge prio sum: %v\n", badgePrioSum)
}

func getPrio(n int) int {
	if n > 96 {
		return n - 96
	}
	return n - 65 + 27
}
