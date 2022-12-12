package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	solution1 := monkeyBusiness(lines, 20, func(monkeys []*Monkey, worryLevel int) int {
		return worryLevel / 3
	})
	solution2 := monkeyBusiness(lines, 10000, func(monkeys []*Monkey, worryLevel int) int {
		return worryLevel % memoizedMonkeyLCM(monkeys)
	})
	fmt.Printf("Solutions: %v, %v\n", solution1, solution2)
}

var memoized int = 0

func memoizedMonkeyLCM(monkeys []*Monkey) int {
	if memoized == 0 {
		divisiors := make([]int, 0)
		for _, monkey := range monkeys {
			divisiors = append(divisiors, monkey.test.divisibleBy)
		}
		memoized = shared.LCM(divisiors...)
	}
	return memoized
}

func monkeyBusiness(lines []string, rounds int, relieveStress func(monkeys []*Monkey, worryLevel int) int) int {
	monkeys := parseMonkeys(lines)
	for round := 1; round < rounds+1; round++ {
		for _, monkey := range monkeys {
			itemsToDelete := make([]int, 0)
			for j, item := range monkey.items {
				variables := make(map[string]int)
				variables["old"] = item
				monkey.operation.apply(variables)
				monkey.inspectionCount++
				item = relieveStress(monkeys, variables["new"])
				var target int
				if monkey.test.evaluate(item) {
					target = monkey.ifTrue.target
				} else {
					target = monkey.ifFalse.target
				}

				monkeys[target].items = append(monkeys[target].items, item)
				itemsToDelete = append(itemsToDelete, j)
			}
			for i := len(itemsToDelete) - 1; i >= 0; i-- {
				idx := itemsToDelete[i]
				monkey.items = append(monkey.items[:idx], monkey.items[idx+1:]...)
			}
		}
	}
	inspections := make([]int, len(monkeys))
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspectionCount)
	}
	sort.Ints(inspections)
	top2 := inspections[len(inspections)-2:]
	return top2[0] * top2[1]
}

type Monkey struct {
	index           int
	inspectionCount int
	items           []int
	operation       Operation
	test            Test
	ifTrue          ThrowToOther
	ifFalse         ThrowToOther
}

func parseMonkeys(lines []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	for _, line := range lines {

		splitLine := strings.Split(line, ":")
		prefix := strings.TrimSpace(splitLine[0])

		if strings.HasPrefix(prefix, "Monkey") {
			index, _ := strconv.Atoi(strings.Split(prefix, " ")[1])

			monkeys = append(monkeys, &Monkey{
				index: index,
				items: make([]int, 0),
			})
		} else if prefix == "Starting items" {
			items := strings.Split(splitLine[1], ",")
			for _, item := range items {
				parsed, _ := strconv.Atoi(strings.TrimSpace(item))
				monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, parsed)
			}
		} else if prefix == "Operation" {
			monkeys[len(monkeys)-1].operation = parseOperation(strings.TrimSpace(splitLine[1]))
		} else if prefix == "Test" {
			monkeys[len(monkeys)-1].test = parseTest(strings.TrimSpace(splitLine[1]))
		} else if prefix == "If true" {
			monkeys[len(monkeys)-1].ifTrue = parseThrowToOther(strings.TrimSpace(splitLine[1]))
		} else if prefix == "If false" {
			monkeys[len(monkeys)-1].ifFalse = parseThrowToOther(strings.TrimSpace(splitLine[1]))
		}
	}
	return monkeys
}

type Test struct {
	divisibleBy int
}

func parseTest(s string) Test {
	rawDivisibleBy := strings.Split(s, " ")[2]
	if divisibleBy, err := strconv.Atoi(rawDivisibleBy); err != nil {
		panic(err)
	} else {
		return Test{
			divisibleBy: divisibleBy,
		}
	}
}

func (t *Test) evaluate(n int) bool {
	return n%t.divisibleBy == 0
}

type Operation struct {
	target    string
	operation string
	a         string
	b         string
}

func parseOperation(o string) Operation {
	s := strings.Split(o, " ")
	target := s[0]
	operation := s[3]
	a := s[2]
	b := s[4]

	return Operation{
		target:    target,
		operation: operation,
		a:         a,
		b:         b,
	}
}

func (o *Operation) evaluate(s string, variables map[string]int) int {
	if v, exists := variables[s]; exists {
		return v
	}
	v, _ := strconv.Atoi(s)
	return v
}

func (o *Operation) apply(variables map[string]int) {
	a := o.evaluate(o.a, variables)
	b := o.evaluate(o.b, variables)

	result := o.evaluateOperation(a, b)

	variables[o.target] = result
}

func (o *Operation) evaluateOperation(a int, b int) int {
	switch o.operation {
	case "+":
		return a + b
	case "*":
		return a * b
	default:
		panic("Unknown operation " + o.operation)
	}
}

type ThrowToOther struct {
	target int
}

func parseThrowToOther(s string) ThrowToOther {
	i, _ := strconv.Atoi(strings.Split(s, " ")[3])
	return ThrowToOther{
		target: i,
	}
}
