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
	solution1 := simulateSandFlow(lines, Point{500, 0}, true)
	solution2 := simulateSandFlow(lines, Point{500, 0}, false)
	fmt.Printf("Solution: %v, %v\n", solution1, solution2)
}

func simulateSandFlow(lines []string, sandDripsFrom Point, hasAbyss bool) int {
	rockMap, normalizedSandDripsFrom := parseLines(lines, sandDripsFrom, hasAbyss)
	printMap(rockMap)
	numSandGrainsAtRest := 0
	for {
		sandGrainCameToRestAt := pourNextGrainOfSand(rockMap, normalizedSandDripsFrom)
		if sandGrainCameToRestAt == nil {
			break
		}
		numSandGrainsAtRest++
		if sandGrainCameToRestAt.x == normalizedSandDripsFrom.x && sandGrainCameToRestAt.y == normalizedSandDripsFrom.y {
			break
		}
	}
	printMap(rockMap)
	return numSandGrainsAtRest
}

func pourNextGrainOfSand(rockMap [][]rune, sandDripsFrom Point) *Point {
	x := sandDripsFrom.x
	y := sandDripsFrom.y
	width := len(rockMap) - 1
	height := len(rockMap[0]) - 1

	if rockMap[x][y] == 'o' {
		printMap(rockMap)
		panic("Should have exited")
	}

	for {
		if y+1 > height {
			return nil
		} else if rockMap[x][y+1] == '.' {
			y += 1
		} else if x-1 < 0 {
			return nil
		} else if rockMap[x-1][y+1] == '.' {
			x -= 1
			y += 1
		} else if x+1 > width {
			return nil
		} else if rockMap[x+1][y+1] == '.' {
			x += 1
			y += 1
		} else {
			rockMap[x][y] = 'o'
			return &Point{x, y}
		}
	}
}

type Point struct {
	x int
	y int
}

func parseLines(lines []string, sandDripsFrom Point, hasAbyss bool) ([][]rune, Point) {
	positions := make(map[int]map[int]bool)

	for _, line := range lines {
		points := make([]Point, 0)
		for _, rawPoint := range strings.Split(line, " -> ") {
			spl := strings.Split(rawPoint, ",")
			x, _ := strconv.Atoi(spl[0])
			y, _ := strconv.Atoi(spl[1])
			points = append(points, Point{x, y})
		}
		for i := 0; i < len(points)-1; i++ {
			p := points[i]
			q := points[i+1]

			drawLine(positions, p, q)
		}
	}

	minX := sandDripsFrom.x
	maxX := sandDripsFrom.x
	minY := sandDripsFrom.y
	maxY := sandDripsFrom.y

	for x := range positions {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		for y := range positions[x] {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	if !hasAbyss {
		// Add anti-abyss line
		minX = minX - height*2
		maxX = maxX + height*2
		maxY = maxY + 2

		p := Point{minX, maxY}
		q := Point{maxX, maxY}
		drawLine(positions, p, q)
	}
	width = maxX - minX + 1
	height = maxY - minY + 1

	rockMap := make([][]rune, width)
	for x := range rockMap {
		rockMap[x] = make([]rune, height)
	}

	for x := range rockMap {
		for y := range rockMap[x] {
			c := '.'
			if x+minX == sandDripsFrom.x && y+minY == sandDripsFrom.y {
				c = '+'
			} else if positions[x+minX][y+minY] {
				c = '#'
			}
			rockMap[x][y] = c
		}
	}

	return rockMap, Point{sandDripsFrom.x - minX, sandDripsFrom.y - minY}
}

func drawLine(positions map[int]map[int]bool, p Point, q Point) {
	// Draw the line from p to q in the positions map
	xDiff := shared.IAbs(p.x - q.x)
	yDiff := shared.IAbs(p.y - q.y)

	if xDiff > 0 && yDiff == 0 {
		var fromX, toX int
		if p.x <= q.x {
			fromX = p.x
			toX = q.x
		} else {
			fromX = q.x
			toX = p.x
		}
		for x := fromX; x <= toX; x++ {
			if _, exists := positions[x]; !exists {
				positions[x] = make(map[int]bool)
			}
			positions[x][p.y] = true
		}
	} else if yDiff > 0 && xDiff == 0 {
		var fromY, toY int
		if p.y <= q.y {
			fromY = p.y
			toY = q.y
		} else {
			fromY = q.y
			toY = p.y
		}
		if _, exists := positions[p.x]; !exists {
			positions[p.x] = make(map[int]bool)
		}
		for y := fromY; y <= toY; y++ {
			positions[p.x][y] = true
		}
	} else {
		panic("Both x and y changed")
	}
}

func printMap(rockMap [][]rune) {
	for y := range rockMap[0] {
		for x := range rockMap {
			fmt.Print(string(rockMap[x][y]))
		}
		fmt.Println()
	}
}
