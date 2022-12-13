package main

import (
	"fmt"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	solution1 := shortestPath(lines, []rune{'S'})
	solution2 := shortestPath(lines, []rune{'S', 'a'})
	fmt.Printf("Solutions: %v, %v\n", solution1, solution2)
}

type Vertex struct {
	index     int
	x         int
	y         int
	height    int
	rawHeight rune
	parent    *Vertex
}

func NewVertex(x int, y int, index int, rawHeight rune) *Vertex {
	v := new(Vertex)
	v.rawHeight = rawHeight
	v.index = index
	v.x = x
	v.y = y
	if rawHeight == 'S' {
		rawHeight = 'a'
	} else if rawHeight == 'E' {
		rawHeight = 'z'
	}
	v.height = int(rawHeight) - int('a')
	return v
}

func shortestPath(lines []string, fromHeights []rune) int {
	width := len(lines[0])
	height := len(lines)
	vertices := make([]*Vertex, width*height)
	// edges[i][j] is true if you can walk from vertices[i] to vertices[j]
	edges := make([][]bool, len(vertices))
	for i := range edges {
		edges[i] = make([]bool, len(vertices))
	}

	for row, line := range lines {
		for col, rawHeight := range line {
			index := row*width + col
			vertices[index] = NewVertex(col, row, index, rawHeight)
		}
	}
	for i := range edges {
		for j := range edges[i] {
			xSteps := shared.IAbs(vertices[i].x - vertices[j].x)
			ySteps := shared.IAbs(vertices[i].y - vertices[j].y)
			isAdjacent := (xSteps == 1 && ySteps == 0) || (xSteps == 0 && ySteps == 1)
			heightDifference := vertices[j].height - vertices[i].height
			hasAcceptableHeightDifference := heightDifference <= 1
			if i != j && isAdjacent && hasAcceptableHeightDifference {
				edges[i][j] = true
			}
		}
	}

	fromVertices := make([]int, 0)
	endAt := 0
	for i, v := range vertices {
		if v.rawHeight == 'E' {
			endAt = i
			continue
		}
		for _, r := range fromHeights {
			if r == v.rawHeight {
				fromVertices = append(fromVertices, i)
			}
		}
	}
	shortest := shared.MaxInt
	for _, startFrom := range fromVertices {
		v := BFS(vertices, edges, startFrom, endAt)
		if v == nil {
			continue
		}
		path := getPath(v)
		printPath(width, height, path)
		steps := countSteps(v)
		if steps < shortest {
			shortest = steps
		}
		for _, v := range vertices {
			v.parent = nil
		}
	}
	return shortest
}

func BFS(vertices []*Vertex, edges [][]bool, startFrom int, endAt int) *Vertex {
	explored := make([]bool, len(vertices))
	q := shared.NewQueue[Vertex]()
	q.Offer(vertices[startFrom])
	explored[startFrom] = true

	for !q.IsEmpty() {
		v, _ := q.Poll()
		i := v.index
		if i == endAt {
			return v
		}
		for j, hasEdge := range edges[i] {
			if hasEdge && !explored[j] {
				explored[j] = true
				w := vertices[j]
				w.parent = v
				q.Offer(w)
			}
		}
	}
	return nil
}

func printExplored(explored []bool, vertices []*Vertex, width int, height int) {
	matrix := make([][]rune, width)
	for i := range matrix {
		matrix[i] = make([]rune, height)
	}

	for i, isExplored := range explored {
		v := vertices[i]
		if isExplored {
			matrix[v.x][v.y] = '#'
		} else {
			matrix[v.x][v.y] = v.rawHeight
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(string(matrix[x][y]))
		}
		fmt.Print("\n")
	}

	fmt.Println("---")
}

func printPath(width int, height int, path []*Vertex) {
	matrix := make([][]rune, width)
	for i := range matrix {
		matrix[i] = make([]rune, height)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	for i, v := range path {
		if i == len(path)-1 {
			// fmt.Printf("Reached goal at (%v,%v)\n", v.x, v.y)
			matrix[v.x][v.y] = 'E'
			continue
		}
		if i == 0 {
			// fmt.Printf("Starting at (%v,%v)\n", v.x, v.y)
		}
		w := path[i+1]
		if w.x > v.x {
			matrix[v.x][v.y] = '>'
		} else if w.y > v.y {
			matrix[v.x][v.y] = 'v'
		} else if w.x < v.x {
			matrix[v.x][v.y] = '<'
		} else if w.y < v.y {
			matrix[v.x][v.y] = '^'
		} else {
			panic("Unreachable?")
		}
		// fmt.Printf("Going %v from (%v,%v) to (%v,%v)\n", string(matrix[v.x][v.y]), v.x, v.y, w.x, w.y)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(string(matrix[x][y]))
		}
		fmt.Print("\n")
	}
}

func getPath(v *Vertex) []*Vertex {
	if v.parent == nil {
		return append(make([]*Vertex, 0), v)
	}

	return append(getPath(v.parent), v)
}

func countSteps(v *Vertex) int {
	if v.parent == nil {
		return 0
	}
	return 1 + countSteps(v.parent)
}
