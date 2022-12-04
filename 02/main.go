package main

import (
	"fmt"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("./input.txt")

	rounds := strings.Split(contents, "\n")
	totalScore := 0
	for _, round := range rounds {
		if round == "" {
			continue
		}
		score := checkRound(round)

		totalScore += score
	}

	fmt.Printf("Score: %v", totalScore)
}

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

func parseElfShape(s string) Shape {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		fmt.Println(s)
		panic("Invalid shape")
	}
}

func parseMyShape(s string) Shape {
	switch s {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		fmt.Println(s)
		panic("Invalid shape")
	}
}

type Outcome int

const (
	Win Outcome = iota
	Lose
	Draw
)

func parseDesiredOutcome(do string) Outcome {
	switch do {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		fmt.Println(do)
		panic("Invalid desired outcome")
	}
}

func getShapeScore(s Shape) int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		fmt.Println(s)
		panic("Invalid shape")
	}
}

func getDesiredShape(elfShape Shape, desiredOutcome Outcome) Shape {
	if desiredOutcome == Draw {
		return elfShape
	} else if desiredOutcome == Win {
		switch elfShape {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	} else { // Lose
		switch elfShape {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}
	panic("Invalid shape or outcome")
}

func checkRound(round string) int {
	chosenShapes := strings.Split(round, " ")
	elfShape := parseElfShape(chosenShapes[0])
	// myShape := parseMyShape(chosenShapes[1])
	desiredOutcome := parseDesiredOutcome(chosenShapes[1])
	myShape := getDesiredShape(elfShape, desiredOutcome)
	myShapeScore := getShapeScore(myShape)

	if myShape == elfShape {
		return myShapeScore + 3
	} else if shapeWins(myShape, elfShape) {
		return myShapeScore + 6
	} else {
		return myShapeScore
	}
}

func shapeWins(s1 Shape, s2 Shape) bool {
	return s1 == Rock && s2 == Scissors || s1 == Paper && s2 == Rock || s1 == Scissors && s2 == Paper
}
