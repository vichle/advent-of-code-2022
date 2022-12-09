package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

type Position struct {
	x int
	y int
}

func main() {
	contents := shared.ReadFileContents("input.txt")
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	numKnots := 10
	visited := make(map[int]map[int]bool)
	visited[0] = make(map[int]bool)
	visited[0][0] = true
	knots := make([]Position, 0)
	for i := 0; i < numKnots; i++ {
		knots = append(knots, Position{x: 0, y: 0})
	}
	uniqueVisited := 1
	for _, line := range lines {
		move := strings.Split(line, " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "U":
				knots[0].y -= 1
				break
			case "D":
				knots[0].y += 1
				break
			case "R":
				knots[0].x += 1
				break
			case "L":
				knots[0].x -= 1
				break
			default:
				panic("Invalid direction " + direction)
			}

			for i := 1; i < len(knots); i++ {
				diffX := shared.IAbs(knots[i-1].x - knots[i].x)
				diffY := shared.IAbs(knots[i-1].y - knots[i].y)

				if diffX > 1 || (diffX == 1 && diffY > 1) {
					knots[i].x = knots[i].x + shared.ISign(knots[i-1].x-knots[i].x)
				}
				if diffY > 1 || (diffY == 1 && diffX > 1) {
					knots[i].y = knots[i].y + shared.ISign(knots[i-1].y-knots[i].y)
				}
			}
			tX := knots[len(knots)-1].x
			tY := knots[len(knots)-1].y
			if _, ok := visited[tX]; !ok {
				visited[tX] = make(map[int]bool)
			}
			if _, ok := visited[tX][tY]; !ok {
				visited[tX][tY] = true
				uniqueVisited += 1
			}
		}
	}

	fmt.Printf("Solution: %v\n", uniqueVisited)
}
