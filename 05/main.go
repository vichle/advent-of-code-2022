package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	rows := strings.Split(contents, "\n")
	state := make([]shared.Stack[string], 0)
	splitAt := -1
	for i, row := range rows {
		if row == "" {
			splitAt = i
			break
		}
	}
	// Create stacks
	for i := 0; i < (len(rows[0])+1)/4; i++ {
		state = append(state, shared.NewStack[string]())
	}

	for i := splitAt - 2; i >= 0; i-- {
		row := " " + rows[i] // Add space at beginning to ease splitting
		for j := 0; j < len(row); j += 4 {
			crate := row[j+2 : j+3]
			if crate != " " {
				state[j/4].Put(crate)
			}
		}

	}
	for _, instructions := range rows[splitAt+1:] {
		if instructions == "" {
			break
		}

		ins := strings.Split(instructions, " ")
		numMoves, _ := strconv.Atoi(ins[1])
		fromStack, _ := strconv.Atoi(ins[3])
		toStack, _ := strconv.Atoi(ins[5])

		fmt.Println("Input", instructions)
		fmt.Printf("Parsed: move %v from %v to %v\n", numMoves, fromStack, toStack)

		// for i := 0; i < numMoves; i++ {
		// 	e := state[fromStack-1].Pop()
		// 	state[toStack-1].Push(e)
		// }

		es := state[fromStack-1].PopMany(numMoves)
		state[toStack-1].PutMany(es)
	}
	ans := ""
	for _, s := range state {
		ans += s.Peek()
	}

	fmt.Printf("Solution: %v\n", ans)
}
