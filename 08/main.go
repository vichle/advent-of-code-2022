package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	size := len(lines)
	trees := make([][]int, size)
	for i, col := range lines {
		trees[i] = make([]int, size)
		for j, treeHeight := range strings.Split(col, "") {
			trees[i][j], _ = strconv.Atoi(treeHeight)
		}
	}
	isVisible := make([][]bool, size)
	score := make([][]int, size)
	for i, _ := range trees {
		isVisible[i] = make([]bool, size)
		score[i] = make([]int, size)
		for j, _ := range trees {
			isVisible[i][j] = checkVisibility(i, j, trees)
			score[i][j] = scenicScore(i, j, trees, false)
			fmt.Printf("[%v,%v] (%v) = %v   ", i, j, trees[i][j], score[i][j])
		}
		fmt.Println("")
	}

	countVisible := 0
	maxScore := 0
	for i, _ := range isVisible {
		for j, visible := range isVisible[i] {
			if visible {
				countVisible += 1
			}
			if score[i][j] > maxScore {
				maxScore = score[i][j]
			}
		}
	}

	scenicScore(3, 2, trees, true)
	fmt.Printf("Num visible: %v\n", countVisible)
	fmt.Printf("Max scenic score: %v\n", maxScore)
}

func checkVisibility(i int, j int, trees [][]int) bool {
	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees)-1 {
		return true
	}
	visibleLeft := true
	visibleRight := true
	visibleUp := true
	visibleDown := true
	// Check left
	for k := j - 1; k >= 0; k-- {
		if trees[i][j] <= trees[i][k] {
			visibleLeft = false
			break
		}
	}
	// Check right
	for k := j + 1; k < len(trees[i]); k++ {
		if trees[i][j] <= trees[i][k] {
			visibleRight = false
			break
		}
	}
	// Check up
	for k := i - 1; k >= 0; k-- {
		if trees[i][j] <= trees[k][j] {
			visibleUp = false
			break
		}
	}
	// Check down
	for k := i + 1; k < len(trees[i]); k++ {
		if trees[i][j] <= trees[k][j] {
			visibleDown = false
			break
		}
	}

	return visibleLeft || visibleRight || visibleUp || visibleDown
}

func scenicScore(i int, j int, trees [][]int, shouldPrint bool) int {
	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees)-1 {
		return 0
	}
	scoreLeft := 0
	scoreRight := 0
	scoreUp := 0
	scoreDown := 0
	// Check left
	for k := j - 1; k >= 0; k-- {
		scoreLeft += 1
		if trees[i][j] <= trees[i][k] {
			break
		}
	}
	// Check right
	for k := j + 1; k < len(trees[i]); k++ {
		scoreRight += 1
		if trees[i][j] <= trees[i][k] {
			break
		}
	}
	// Check up
	for k := i - 1; k >= 0; k-- {
		scoreUp += 1
		if trees[i][j] <= trees[k][j] {
			break
		}
	}
	// Check down
	for k := i + 1; k < len(trees[i]); k++ {
		scoreDown += 1
		if trees[i][j] <= trees[k][j] {
			break
		}
	}

	if shouldPrint {
		fmt.Println(scoreLeft, scoreRight, scoreDown, scoreUp)
	}

	return scoreLeft * scoreRight * scoreDown * scoreUp
}
